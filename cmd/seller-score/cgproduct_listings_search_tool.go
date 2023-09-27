package seller_score

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"

	"github.com/Longoria-Yu/test-scripts/common"
	contentfeed_proto "github.com/Longoria-Yu/test-scripts/pb/content-feed"
	listing_proto "github.com/Longoria-Yu/test-scripts/pb/listing"
)

type ExpGroup int

const (
	ControlGroup    ExpGroup = 0
	TreatmentAGroup ExpGroup = 1
	TreatmentBGroup ExpGroup = 2
)

type SellerScoreTestCase struct {
	Idx              int
	CGProductID      string
	CGProductName    string
	LayeredCondition string
	Color            string
	Storage          string
}

type CSVRow struct {
	ListingID    int64
	SellerID     string
	SellerScore  string
	ControlGroup *ScoreInfo
	TreatmentA   *ScoreInfo
	TreatmentB   *ScoreInfo
}

type ScoreInfo struct {
	ListingScore string
	Probability  string
	SellerWeight string
}

func CallSearchCGProductListingsV10(c contentfeed_proto.ContentFeedServiceClient) {
	fmt.Println("making grpc SearchCGProductListingsV10 call")
	var (
		count     int32 = 50
		countryId int64 = common.Singapore
	)

	// Get the Test Cases
	log.Printf("Getting the tes cases from the CSV\n")
	testCases := getSellerScoreTestCaseFromCSV()
	log.Printf("Retrieving the tes cases from the CSV successfully\n")

	//var tcWC sync.WaitGroup
	for tcIdx, tc := range testCases {
		//tcWC.Add(1)
		tc.Idx = tcIdx + 1
		log.Printf("Starting TestCase: %d, Test Case Info: %+v\n", tc.Idx, tc)

		//var wc sync.WaitGroup
		listingScoreInfo := make(map[int64]*CSVRow)
		//	for userIdx, userId := range []int64{12345, 35275645, 11620696} {
		for userIdx, userId := range []int64{123, 34488668, 23644165} {
			//wc.Add(1)
			log.Printf("Starting TestCase: %d, userIdx: %d userId: %d\n", tc.Idx, userIdx, userId)

			func(tc SellerScoreTestCase, userIdx int, userId int64) {
				req := &contentfeed_proto.SearchCGProductListingsRequestV1{
					Count:     count,
					CountryId: countryId,
					Filters: []*contentfeed_proto.FilterParamV4{
						{
							FieldName: "cgproduct_id",
							FilterType: &contentfeed_proto.FilterParamV4_IdsOrKeywords{
								IdsOrKeywords: &contentfeed_proto.IdsOrKeywords{
									Value: []string{tc.CGProductID},
								},
							},
						},
						{
							FieldName: "status",
							FilterType: &contentfeed_proto.FilterParamV4_IdsOrKeywords{
								IdsOrKeywords: &contentfeed_proto.IdsOrKeywords{
									Value: []string{"L"},
								},
							},
						},
						{
							FieldName: "is_advance_promise_enabled",
							FilterType: &contentfeed_proto.FilterParamV4_IdsOrKeywords{
								IdsOrKeywords: &contentfeed_proto.IdsOrKeywords{
									Value: []string{"true"},
								},
							},
						},
						{
							FieldName: "layered_condition",
							FilterType: &contentfeed_proto.FilterParamV4_IdsOrKeywords{
								IdsOrKeywords: &contentfeed_proto.IdsOrKeywords{
									Value: []string{tc.LayeredCondition},
								},
							},
						},
						{
							FieldName: "_cgproduct.color",
							FilterType: &contentfeed_proto.FilterParamV4_IdsOrKeywords{
								IdsOrKeywords: &contentfeed_proto.IdsOrKeywords{
									Value: []string{tc.Color},
								},
							},
						},
						{
							FieldName: "_cgproduct.storage",
							FilterType: &contentfeed_proto.FilterParamV4_IdsOrKeywords{
								IdsOrKeywords: &contentfeed_proto.IdsOrKeywords{
									Value: []string{tc.Storage},
								},
							},
						},
					},
					Sort: contentfeed_proto.SearchCGProductListingsRequestV1_SORT_TYPE_RECOMMENDED,
					AuthUserContainer: &contentfeed_proto.AuthUserContainer{
						AuthUserId:  userId,
						CountryId:   common.Singapore,
						CountryCode: common.SingaporeCode,
						IsAdmin:     true,
					},
					IncludeDebuggingInfo: true,
				}

				ctx := context.Background()
				ctx = metadata.AppendToOutgoingContext(ctx,
					"build-no", "999999",
					"platform", "ios",
					"accept-language", "en",
				)

				var (
					resp *contentfeed_proto.SearchCGProductListingsResponseV1
					err  error
				)

				for i := 0; i <= 2; i++ {
					if resp == nil {
						log.Printf("calling SearchCGProductListingsV10 for testCase: %d with the userId: %d\n", tc.Idx, userId)
						resp, err = c.SearchCGProductListingsV10(ctx, req)
					}

					if err != nil {
						log.Printf("failed calling SearchCGProductListingsV10 for testCase: %d with the userId: %d\n", tc.Idx, userId)
						if i <= 2 {
							time.Sleep(1 * time.Second)
						} else {
							log.Fatalf("Failed to call SearchCGProductListingsV10: %v", err)
						}
					} else {
						for _, listing := range resp.CgproductListing {
							updateScoreInfoMap(listingScoreInfo, ExpGroup(userIdx), listing)
						}

						//wc.Done()
						break
					}
				}
			}(tc, userIdx, userId)
		}

		//wc.Wait()
		createResultFile(tc, listingScoreInfo)
		//tcWC.Done()
	}
	//tcWC.Wait()
	log.Println("Running the script successfully!")
}

func updateScoreInfoMap(listingScoreInfo map[int64]*CSVRow, expGroup ExpGroup, listing *contentfeed_proto.SearchCGProductListingsResponseV1_CGProductListing) {
	lsi, ok := listingScoreInfo[listing.GetLegacyId()]
	if !ok {
		lsi = &CSVRow{
			ListingID:    listing.GetLegacyId(),
			ControlGroup: &ScoreInfo{},
			TreatmentA:   &ScoreInfo{},
			TreatmentB:   &ScoreInfo{},
		}
		listingScoreInfo[listing.GetLegacyId()] = lsi
	}

	strs := strings.Split(listing.GetExplain(), "pdp")
	infos := strings.FieldsFunc(strs[len(strs)-1], func(r rune) bool {
		return r == '[' || r == ']' || r == ','
	})

	var row *ScoreInfo
	switch expGroup {
	case ControlGroup:
		row = lsi.ControlGroup
	case TreatmentAGroup:
		row = lsi.TreatmentA
	case TreatmentBGroup:
		row = lsi.TreatmentB
	}

	for i, info := range infos {
		if !strings.Contains(info, "seller-factors") {
			data := strings.FieldsFunc(info, func(r rune) bool {
				return r == ':' || r == ' '
			})[1]

			switch i {
			case 0:
				row.ListingScore = data
			case 1:
				row.Probability = data
			case 3:
				lsi.SellerID = data
			case 4:
				lsi.SellerScore = data
			case 5:
				row.SellerWeight = data
			}
		}
	}
}

func createResultFile(tc SellerScoreTestCase, listingScoreInfo map[int64]*CSVRow) {
	// Create or open the CSV file for writing
	fileName := fmt.Sprintf("./%s_%s_%s_%s.csv", tc.CGProductName, tc.Color, tc.Storage, getLayeredCondition(tc.LayeredCondition))
	log.Printf("creating the CSV file: %s", fileName)

	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Initialize a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header row
	header := []string{"Index", "Listing ID", "Seller ID", "Seller Score",
		"Control_Total_Score", "Control_Probability", "Control_Seller_Weight",
		"TreatmentA_Total_Score", "TreatmentA_Probability", "TreatmentA_Seller_Weight",
		"TreatmentB_Total_Score", "TreatmentB_Probability", "TreatmentB_Seller_Weight",
	}
	if err := writer.Write(header); err != nil {
		log.Fatal(err)
	}

	rowIdx := 1
	for listingID, info := range listingScoreInfo {
		row := []string{strconv.Itoa(rowIdx), strconv.FormatInt(listingID, 10), info.SellerID, info.SellerScore,
			info.ControlGroup.ListingScore, info.ControlGroup.Probability, info.ControlGroup.SellerWeight,
			info.TreatmentA.ListingScore, info.TreatmentA.Probability, info.TreatmentA.SellerWeight,
			info.TreatmentB.ListingScore, info.TreatmentB.Probability, info.TreatmentB.SellerWeight,
		}

		if err := writer.Write(row); err != nil {
			log.Fatal(err)
		}
		rowIdx += 1
	}
	log.Printf("completing the CSV file: %s", fileName)
}

func getLayeredCondition(lc string) string {
	lcInt, _ := strconv.ParseInt(lc, 10, 32)
	condition := listing_proto.Condition(lcInt)
	switch condition {
	case listing_proto.Condition_CONDITION_BRAND_NEW:
		return "Brand new"
	case listing_proto.Condition_CONDITION_LIKE_NEW:
		return "Like new"
	case listing_proto.Condition_CONDITION_LIGHTLY_USED:
		return "Lightly used"
	case listing_proto.Condition_CONDITION_WELL_USED:
		return "Well used"
	}
	return "Unknown condition"
}

func getSellerScoreTestCaseFromCSV() []SellerScoreTestCase {
	// Open the CSV file for reading
	file, err := os.Open("./seller_score_test_cases.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Initialize a CSV reader
	reader := csv.NewReader(file)

	// Read and process log entries
	var testCases []SellerScoreTestCase
	rowNum := 1
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Reading csv error: %+v\n", err)
			continue
		}

		tc := SellerScoreTestCase{
			Idx:              rowNum,
			CGProductID:      record[0],
			CGProductName:    record[1],
			LayeredCondition: record[2],
			Color:            record[3],
			Storage:          record[4],
		}
		testCases = append(testCases, tc)
		rowNum += 1
	}

	return testCases[1:]
}
