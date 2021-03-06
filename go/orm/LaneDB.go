// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gonggantt/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Lane_sql sql.NullBool
var dummy_Lane_time time.Duration
var dummy_Lane_sort sort.Float64Slice

// LaneAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model laneAPI
type LaneAPI struct {
	gorm.Model

	models.Lane

	// encoding of pointers
	LanePointersEnconding
}

// LanePointersEnconding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type LanePointersEnconding struct {
	// insertion for pointer fields encoding declaration

	// Implementation of a reverse ID for field Gantt{}.Lanes []*Lane
	Gantt_LanesDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Gantt_LanesDBID_Index sql.NullInt64

	// Implementation of a reverse ID for field Group{}.GroupLanes []*Lane
	Group_GroupLanesDBID sql.NullInt64

	// implementation of the index of the withing the slice
	Group_GroupLanesDBID_Index sql.NullInt64
}

// LaneDB describes a lane in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model laneDB
type LaneDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field laneDB.Name
	Name_Data sql.NullString

	// Declation for basic field laneDB.Order
	Order_Data sql.NullInt64
	// encoding of pointers
	LanePointersEnconding
}

// LaneDBs arrays laneDBs
// swagger:response laneDBsResponse
type LaneDBs []LaneDB

// LaneDBResponse provides response
// swagger:response laneDBResponse
type LaneDBResponse struct {
	LaneDB
}

// LaneWOP is a Lane without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type LaneWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Order int `xlsx:"2"`
	// insertion for WOP pointer fields
}

var Lane_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Order",
}

type BackRepoLaneStruct struct {
	// stores LaneDB according to their gorm ID
	Map_LaneDBID_LaneDB *map[uint]*LaneDB

	// stores LaneDB ID according to Lane address
	Map_LanePtr_LaneDBID *map[*models.Lane]uint

	// stores Lane according to their gorm ID
	Map_LaneDBID_LanePtr *map[uint]*models.Lane

	db *gorm.DB
}

func (backRepoLane *BackRepoLaneStruct) GetDB() *gorm.DB {
	return backRepoLane.db
}

// GetLaneDBFromLanePtr is a handy function to access the back repo instance from the stage instance
func (backRepoLane *BackRepoLaneStruct) GetLaneDBFromLanePtr(lane *models.Lane) (laneDB *LaneDB) {
	id := (*backRepoLane.Map_LanePtr_LaneDBID)[lane]
	laneDB = (*backRepoLane.Map_LaneDBID_LaneDB)[id]
	return
}

// BackRepoLane.Init set up the BackRepo of the Lane
func (backRepoLane *BackRepoLaneStruct) Init(db *gorm.DB) (Error error) {

	if backRepoLane.Map_LaneDBID_LanePtr != nil {
		err := errors.New("In Init, backRepoLane.Map_LaneDBID_LanePtr should be nil")
		return err
	}

	if backRepoLane.Map_LaneDBID_LaneDB != nil {
		err := errors.New("In Init, backRepoLane.Map_LaneDBID_LaneDB should be nil")
		return err
	}

	if backRepoLane.Map_LanePtr_LaneDBID != nil {
		err := errors.New("In Init, backRepoLane.Map_LanePtr_LaneDBID should be nil")
		return err
	}

	tmp := make(map[uint]*models.Lane, 0)
	backRepoLane.Map_LaneDBID_LanePtr = &tmp

	tmpDB := make(map[uint]*LaneDB, 0)
	backRepoLane.Map_LaneDBID_LaneDB = &tmpDB

	tmpID := make(map[*models.Lane]uint, 0)
	backRepoLane.Map_LanePtr_LaneDBID = &tmpID

	backRepoLane.db = db
	return
}

// BackRepoLane.CommitPhaseOne commits all staged instances of Lane to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLane *BackRepoLaneStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for lane := range stage.Lanes {
		backRepoLane.CommitPhaseOneInstance(lane)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, lane := range *backRepoLane.Map_LaneDBID_LanePtr {
		if _, ok := stage.Lanes[lane]; !ok {
			backRepoLane.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoLane.CommitDeleteInstance commits deletion of Lane to the BackRepo
func (backRepoLane *BackRepoLaneStruct) CommitDeleteInstance(id uint) (Error error) {

	lane := (*backRepoLane.Map_LaneDBID_LanePtr)[id]

	// lane is not staged anymore, remove laneDB
	laneDB := (*backRepoLane.Map_LaneDBID_LaneDB)[id]
	query := backRepoLane.db.Unscoped().Delete(&laneDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	delete((*backRepoLane.Map_LanePtr_LaneDBID), lane)
	delete((*backRepoLane.Map_LaneDBID_LanePtr), id)
	delete((*backRepoLane.Map_LaneDBID_LaneDB), id)

	return
}

// BackRepoLane.CommitPhaseOneInstance commits lane staged instances of Lane to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoLane *BackRepoLaneStruct) CommitPhaseOneInstance(lane *models.Lane) (Error error) {

	// check if the lane is not commited yet
	if _, ok := (*backRepoLane.Map_LanePtr_LaneDBID)[lane]; ok {
		return
	}

	// initiate lane
	var laneDB LaneDB
	laneDB.CopyBasicFieldsFromLane(lane)

	query := backRepoLane.db.Create(&laneDB)
	if query.Error != nil {
		return query.Error
	}

	// update stores
	(*backRepoLane.Map_LanePtr_LaneDBID)[lane] = laneDB.ID
	(*backRepoLane.Map_LaneDBID_LanePtr)[laneDB.ID] = lane
	(*backRepoLane.Map_LaneDBID_LaneDB)[laneDB.ID] = &laneDB

	return
}

// BackRepoLane.CommitPhaseTwo commits all staged instances of Lane to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLane *BackRepoLaneStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, lane := range *backRepoLane.Map_LaneDBID_LanePtr {
		backRepoLane.CommitPhaseTwoInstance(backRepo, idx, lane)
	}

	return
}

// BackRepoLane.CommitPhaseTwoInstance commits {{structname }} of models.Lane to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLane *BackRepoLaneStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, lane *models.Lane) (Error error) {

	// fetch matching laneDB
	if laneDB, ok := (*backRepoLane.Map_LaneDBID_LaneDB)[idx]; ok {

		laneDB.CopyBasicFieldsFromLane(lane)

		// insertion point for translating pointers encodings into actual pointers
		// This loop encodes the slice of pointers lane.Bars into the back repo.
		// Each back repo instance at the end of the association encode the ID of the association start
		// into a dedicated field for coding the association. The back repo instance is then saved to the db
		for idx, barAssocEnd := range lane.Bars {

			// get the back repo instance at the association end
			barAssocEnd_DB :=
				backRepo.BackRepoBar.GetBarDBFromBarPtr(barAssocEnd)

			// encode reverse pointer in the association end back repo instance
			barAssocEnd_DB.Lane_BarsDBID.Int64 = int64(laneDB.ID)
			barAssocEnd_DB.Lane_BarsDBID.Valid = true
			barAssocEnd_DB.Lane_BarsDBID_Index.Int64 = int64(idx)
			barAssocEnd_DB.Lane_BarsDBID_Index.Valid = true
			if q := backRepoLane.db.Save(barAssocEnd_DB); q.Error != nil {
				return q.Error
			}
		}

		query := backRepoLane.db.Save(&laneDB)
		if query.Error != nil {
			return query.Error
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Lane intance %s", lane.Name))
		return err
	}

	return
}

// BackRepoLane.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for pahse two)
//
func (backRepoLane *BackRepoLaneStruct) CheckoutPhaseOne() (Error error) {

	laneDBArray := make([]LaneDB, 0)
	query := backRepoLane.db.Find(&laneDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	laneInstancesToBeRemovedFromTheStage := make(map[*models.Lane]any)
	for key, value := range models.Stage.Lanes {
		laneInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, laneDB := range laneDBArray {
		backRepoLane.CheckoutPhaseOneInstance(&laneDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		lane, ok := (*backRepoLane.Map_LaneDBID_LanePtr)[laneDB.ID]
		if ok {
			delete(laneInstancesToBeRemovedFromTheStage, lane)
		}
	}

	// remove from stage and back repo's 3 maps all lanes that are not in the checkout
	for lane := range laneInstancesToBeRemovedFromTheStage {
		lane.Unstage()

		// remove instance from the back repo 3 maps
		laneID := (*backRepoLane.Map_LanePtr_LaneDBID)[lane]
		delete((*backRepoLane.Map_LanePtr_LaneDBID), lane)
		delete((*backRepoLane.Map_LaneDBID_LaneDB), laneID)
		delete((*backRepoLane.Map_LaneDBID_LanePtr), laneID)
	}

	return
}

// CheckoutPhaseOneInstance takes a laneDB that has been found in the DB, updates the backRepo and stages the
// models version of the laneDB
func (backRepoLane *BackRepoLaneStruct) CheckoutPhaseOneInstance(laneDB *LaneDB) (Error error) {

	lane, ok := (*backRepoLane.Map_LaneDBID_LanePtr)[laneDB.ID]
	if !ok {
		lane = new(models.Lane)

		(*backRepoLane.Map_LaneDBID_LanePtr)[laneDB.ID] = lane
		(*backRepoLane.Map_LanePtr_LaneDBID)[lane] = laneDB.ID

		// append model store with the new element
		lane.Name = laneDB.Name_Data.String
		lane.Stage()
	}
	laneDB.CopyBasicFieldsToLane(lane)

	// preserve pointer to laneDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_LaneDBID_LaneDB)[laneDB hold variable pointers
	laneDB_Data := *laneDB
	preservedPtrToLane := &laneDB_Data
	(*backRepoLane.Map_LaneDBID_LaneDB)[laneDB.ID] = preservedPtrToLane

	return
}

// BackRepoLane.CheckoutPhaseTwo Checkouts all staged instances of Lane to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLane *BackRepoLaneStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, laneDB := range *backRepoLane.Map_LaneDBID_LaneDB {
		backRepoLane.CheckoutPhaseTwoInstance(backRepo, laneDB)
	}
	return
}

// BackRepoLane.CheckoutPhaseTwoInstance Checkouts staged instances of Lane to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoLane *BackRepoLaneStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, laneDB *LaneDB) (Error error) {

	lane := (*backRepoLane.Map_LaneDBID_LanePtr)[laneDB.ID]
	_ = lane // sometimes, there is no code generated. This lines voids the "unused variable" compilation error

	// insertion point for checkout of pointer encoding
	// This loop redeem lane.Bars in the stage from the encode in the back repo
	// It parses all BarDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	lane.Bars = lane.Bars[:0]
	// 2. loop all instances in the type in the association end
	for _, barDB_AssocEnd := range *backRepo.BackRepoBar.Map_BarDBID_BarDB {
		// 3. Does the ID encoding at the end and the ID at the start matches ?
		if barDB_AssocEnd.Lane_BarsDBID.Int64 == int64(laneDB.ID) {
			// 4. fetch the associated instance in the stage
			bar_AssocEnd := (*backRepo.BackRepoBar.Map_BarDBID_BarPtr)[barDB_AssocEnd.ID]
			// 5. append it the association slice
			lane.Bars = append(lane.Bars, bar_AssocEnd)
		}
	}

	// sort the array according to the order
	sort.Slice(lane.Bars, func(i, j int) bool {
		barDB_i_ID := (*backRepo.BackRepoBar.Map_BarPtr_BarDBID)[lane.Bars[i]]
		barDB_j_ID := (*backRepo.BackRepoBar.Map_BarPtr_BarDBID)[lane.Bars[j]]

		barDB_i := (*backRepo.BackRepoBar.Map_BarDBID_BarDB)[barDB_i_ID]
		barDB_j := (*backRepo.BackRepoBar.Map_BarDBID_BarDB)[barDB_j_ID]

		return barDB_i.Lane_BarsDBID_Index.Int64 < barDB_j.Lane_BarsDBID_Index.Int64
	})

	return
}

// CommitLane allows commit of a single lane (if already staged)
func (backRepo *BackRepoStruct) CommitLane(lane *models.Lane) {
	backRepo.BackRepoLane.CommitPhaseOneInstance(lane)
	if id, ok := (*backRepo.BackRepoLane.Map_LanePtr_LaneDBID)[lane]; ok {
		backRepo.BackRepoLane.CommitPhaseTwoInstance(backRepo, id, lane)
	}
}

// CommitLane allows checkout of a single lane (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutLane(lane *models.Lane) {
	// check if the lane is staged
	if _, ok := (*backRepo.BackRepoLane.Map_LanePtr_LaneDBID)[lane]; ok {

		if id, ok := (*backRepo.BackRepoLane.Map_LanePtr_LaneDBID)[lane]; ok {
			var laneDB LaneDB
			laneDB.ID = id

			if err := backRepo.BackRepoLane.db.First(&laneDB, id).Error; err != nil {
				log.Panicln("CheckoutLane : Problem with getting object with id:", id)
			}
			backRepo.BackRepoLane.CheckoutPhaseOneInstance(&laneDB)
			backRepo.BackRepoLane.CheckoutPhaseTwoInstance(backRepo, &laneDB)
		}
	}
}

// CopyBasicFieldsFromLane
func (laneDB *LaneDB) CopyBasicFieldsFromLane(lane *models.Lane) {
	// insertion point for fields commit

	laneDB.Name_Data.String = lane.Name
	laneDB.Name_Data.Valid = true

	laneDB.Order_Data.Int64 = int64(lane.Order)
	laneDB.Order_Data.Valid = true
}

// CopyBasicFieldsFromLaneWOP
func (laneDB *LaneDB) CopyBasicFieldsFromLaneWOP(lane *LaneWOP) {
	// insertion point for fields commit

	laneDB.Name_Data.String = lane.Name
	laneDB.Name_Data.Valid = true

	laneDB.Order_Data.Int64 = int64(lane.Order)
	laneDB.Order_Data.Valid = true
}

// CopyBasicFieldsToLane
func (laneDB *LaneDB) CopyBasicFieldsToLane(lane *models.Lane) {
	// insertion point for checkout of basic fields (back repo to stage)
	lane.Name = laneDB.Name_Data.String
	lane.Order = int(laneDB.Order_Data.Int64)
}

// CopyBasicFieldsToLaneWOP
func (laneDB *LaneDB) CopyBasicFieldsToLaneWOP(lane *LaneWOP) {
	lane.ID = int(laneDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	lane.Name = laneDB.Name_Data.String
	lane.Order = int(laneDB.Order_Data.Int64)
}

// Backup generates a json file from a slice of all LaneDB instances in the backrepo
func (backRepoLane *BackRepoLaneStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "LaneDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LaneDB, 0)
	for _, laneDB := range *backRepoLane.Map_LaneDBID_LaneDB {
		forBackup = append(forBackup, laneDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Panic("Cannot json Lane ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Panic("Cannot write the json Lane file", err.Error())
	}
}

// Backup generates a json file from a slice of all LaneDB instances in the backrepo
func (backRepoLane *BackRepoLaneStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*LaneDB, 0)
	for _, laneDB := range *backRepoLane.Map_LaneDBID_LaneDB {
		forBackup = append(forBackup, laneDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Lane")
	if err != nil {
		log.Panic("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Lane_Fields, -1)
	for _, laneDB := range forBackup {

		var laneWOP LaneWOP
		laneDB.CopyBasicFieldsToLaneWOP(&laneWOP)

		row := sh.AddRow()
		row.WriteStruct(&laneWOP, -1)
	}
}

// RestoreXL from the "Lane" sheet all LaneDB instances
func (backRepoLane *BackRepoLaneStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoLaneid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Lane"]
	_ = sh
	if !ok {
		log.Panic(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoLane.rowVisitorLane)
	if err != nil {
		log.Panic("Err=", err)
	}
}

func (backRepoLane *BackRepoLaneStruct) rowVisitorLane(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var laneWOP LaneWOP
		row.ReadStruct(&laneWOP)

		// add the unmarshalled struct to the stage
		laneDB := new(LaneDB)
		laneDB.CopyBasicFieldsFromLaneWOP(&laneWOP)

		laneDB_ID_atBackupTime := laneDB.ID
		laneDB.ID = 0
		query := backRepoLane.db.Create(laneDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoLane.Map_LaneDBID_LaneDB)[laneDB.ID] = laneDB
		BackRepoLaneid_atBckpTime_newID[laneDB_ID_atBackupTime] = laneDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "LaneDB.json" in dirPath that stores an array
// of LaneDB and stores it in the database
// the map BackRepoLaneid_atBckpTime_newID is updated accordingly
func (backRepoLane *BackRepoLaneStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoLaneid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "LaneDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Panic("Cannot restore/open the json Lane file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*LaneDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_LaneDBID_LaneDB
	for _, laneDB := range forRestore {

		laneDB_ID_atBackupTime := laneDB.ID
		laneDB.ID = 0
		query := backRepoLane.db.Create(laneDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
		(*backRepoLane.Map_LaneDBID_LaneDB)[laneDB.ID] = laneDB
		BackRepoLaneid_atBckpTime_newID[laneDB_ID_atBackupTime] = laneDB.ID
	}

	if err != nil {
		log.Panic("Cannot restore/unmarshall json Lane file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Lane>id_atBckpTime_newID
// to compute new index
func (backRepoLane *BackRepoLaneStruct) RestorePhaseTwo() {

	for _, laneDB := range *backRepoLane.Map_LaneDBID_LaneDB {

		// next line of code is to avert unused variable compilation error
		_ = laneDB

		// insertion point for reindexing pointers encoding
		// This reindex lane.Lanes
		if laneDB.Gantt_LanesDBID.Int64 != 0 {
			laneDB.Gantt_LanesDBID.Int64 =
				int64(BackRepoGanttid_atBckpTime_newID[uint(laneDB.Gantt_LanesDBID.Int64)])
		}

		// This reindex lane.GroupLanes
		if laneDB.Group_GroupLanesDBID.Int64 != 0 {
			laneDB.Group_GroupLanesDBID.Int64 =
				int64(BackRepoGroupid_atBckpTime_newID[uint(laneDB.Group_GroupLanesDBID.Int64)])
		}

		// update databse with new index encoding
		query := backRepoLane.db.Model(laneDB).Updates(*laneDB)
		if query.Error != nil {
			log.Panic(query.Error)
		}
	}

}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoLaneid_atBckpTime_newID map[uint]uint
