package models

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var dummy_strconv_import strconv.NumError

// ParseAstFile Parse pathToFile and stages all instances
// declared in the file
func ParseAstFile(pathToFile string) {

	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		log.Panic("Path does not exist %s ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	startParser := time.Now()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		log.Panic("Unable to parser ", errParser.Error())
	}

	// astCoordinate := "File "
	// log.Println(// astCoordinate)
	for _, decl := range inFile.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			funcDecl := decl
			// astCoordinate := // astCoordinate + "\tFunction " + funcDecl.Name.Name
			if name := funcDecl.Name; name != nil {
				isOfInterest := strings.Contains(funcDecl.Name.Name, "Injection")
				if !isOfInterest {
					continue
				}
				// log.Println(// astCoordinate)
			}
			if doc := funcDecl.Doc; doc != nil {
				// astCoordinate := // astCoordinate + "\tDoc"
				for _, comment := range doc.List {
					_ = comment
					// astCoordinate := // astCoordinate + "\tComment: " + comment.Text
					// log.Println(// astCoordinate)
				}
			}
			if body := funcDecl.Body; body != nil {
				// astCoordinate := // astCoordinate + "\tBody: "
				for _, stmt := range body.List {
					switch stmt := stmt.(type) {
					case *ast.ExprStmt:
						exprStmt := stmt
						// astCoordinate := // astCoordinate + "\tExprStmt: "
						switch expr := exprStmt.X.(type) {
						case *ast.CallExpr:
							// astCoordinate := // astCoordinate + "\tCallExpr: "
							callExpr := expr
							switch fun := callExpr.Fun.(type) {
							case *ast.Ident:
								ident := fun
								_ = ident
								// astCoordinate := // astCoordinate + "\tIdent: " + ident.Name
								// log.Println(// astCoordinate)
							}
						}
					case *ast.AssignStmt:
						astCoordinate := "\tAssignStmt: "
						// log.Println(// astCoordinate)
						assignStmt := stmt
						instance, id, gongstruct, fieldName := UnmarshallGongstructStaging(assignStmt, astCoordinate)
						_ = instance
						_ = id
						_ = gongstruct
						_ = fieldName
					}
				}
			}
		case *ast.GenDecl:
			genDecl := decl
			// log.Println("\tAST GenDecl: ")
			if doc := genDecl.Doc; doc != nil {
				for _, comment := range doc.List {
					_ = comment
					// log.Println("\tAST Comment: ", comment.Text)
				}
			}
			for _, spec := range genDecl.Specs {
				switch spec := spec.(type) {
				case *ast.ImportSpec:
					importSpec := spec
					if path := importSpec.Path; path != nil {
						// log.Println("\t\tAST Path: ", path.Value)
					}
				}
			}
		}

	}
}

var __gong__map_Indentifiers_gongstructName = make(map[string]string)

// insertion point for identifiers maps
var __gong__map_Arrow = make(map[string]*Arrow)
var __gong__map_Bar = make(map[string]*Bar)
var __gong__map_Gantt = make(map[string]*Gantt)
var __gong__map_Group = make(map[string]*Group)
var __gong__map_Lane = make(map[string]*Lane)
var __gong__map_LaneUse = make(map[string]*LaneUse)
var __gong__map_Milestone = make(map[string]*Milestone)

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging(assignStmt *ast.AssignStmt, astCoordinate_ string) (
	instance any,
	identifier string,
	gongstructName string,
	fieldName string) {
	astCoordinate := "\tAssignStmt: "
	for rank, expr := range assignStmt.Lhs {
		if rank > 0 {
			continue
		}
		switch expr := expr.(type) {
		case *ast.Ident:
			// we are on a variable declaration
			ident := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + ident.Name
			// log.Println(astCoordinate)
			identifier = ident.Name
		case *ast.SelectorExpr:
			// we are on a variable assignement
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tLhs" + "." + selectorExpr.X.(*ast.Ident).Name + "." + selectorExpr.Sel.Name
			// log.Println(astCoordinate)
			identifier = selectorExpr.X.(*ast.Ident).Name
			fieldName = selectorExpr.Sel.Name
		}
	}
	for _, expr := range assignStmt.Rhs {
		// astCoordinate := astCoordinate + "\tRhs"
		switch expr := expr.(type) {
		case *ast.CallExpr:
			callExpr := expr
			// astCoordinate := astCoordinate + "\tFun"
			switch fun := callExpr.Fun.(type) {
			// the is Fun      Expr
			// function expression xxx.Stage()
			case *ast.SelectorExpr:
				selectorExpr := fun
				// astCoordinate := astCoordinate + "\tSelectorExpr"
				switch x := selectorExpr.X.(type) {
				case *ast.ParenExpr:
					// A ParenExpr node represents a parenthesized expression.
					// the is the
					//   { Name : "A1"}
					// astCoordinate := astCoordinate + "\tX"
					parenExpr := x
					switch x := parenExpr.X.(type) {
					case *ast.UnaryExpr:
						unaryExpr := x
						// astCoordinate := astCoordinate + "\tUnaryExpr"
						switch x := unaryExpr.X.(type) {
						case *ast.CompositeLit:
							instanceName := "NoName yet"
							compositeLit := x
							// astCoordinate := astCoordinate + "\tX(CompositeLit)"
							for _, elt := range compositeLit.Elts {
								// astCoordinate := astCoordinate + "\tElt"
								switch elt := elt.(type) {
								case *ast.KeyValueExpr:
									// This is expression
									//     Name: "A1"
									keyValueExpr := elt
									// astCoordinate := astCoordinate + "\tKeyValueExpr"
									switch key := keyValueExpr.Key.(type) {
									case *ast.Ident:
										ident := key
										_ = ident
										// astCoordinate := astCoordinate + "\tKey" + "." + ident.Name
										// log.Println(astCoordinate)
									}
									switch value := keyValueExpr.Value.(type) {
									case *ast.BasicLit:
										basicLit := value
										// astCoordinate := astCoordinate + "\tBasicLit Value" + "." + basicLit.Value
										// log.Println(astCoordinate)
										instanceName = basicLit.Value

										// remove first and last char
										instanceName = instanceName[1 : len(instanceName)-1]
									}
								}
							}
							astCoordinate2 := astCoordinate + "\tType"
							_ = astCoordinate2
							switch type_ := compositeLit.Type.(type) {
							case *ast.SelectorExpr:
								slcExpr := type_
								// astCoordinate := astCoordinate2 + "\tSelectorExpr"
								switch X := slcExpr.X.(type) {
								case *ast.Ident:
									ident := X
									_ = ident
									// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
									// log.Println(astCoordinate)
								}
								if Sel := slcExpr.Sel; Sel != nil {
									// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
									// log.Println(astCoordinate)

									gongstructName = Sel.Name
									// this is the place where an instance is created
									switch gongstructName {
									// insertion point for identifiers
									case "Arrow":
										instanceArrow := (&Arrow{Name: instanceName}).Stage()
										instance = any(instanceArrow)
										__gong__map_Arrow[identifier] = instanceArrow
									case "Bar":
										instanceBar := (&Bar{Name: instanceName}).Stage()
										instance = any(instanceBar)
										__gong__map_Bar[identifier] = instanceBar
									case "Gantt":
										instanceGantt := (&Gantt{Name: instanceName}).Stage()
										instance = any(instanceGantt)
										__gong__map_Gantt[identifier] = instanceGantt
									case "Group":
										instanceGroup := (&Group{Name: instanceName}).Stage()
										instance = any(instanceGroup)
										__gong__map_Group[identifier] = instanceGroup
									case "Lane":
										instanceLane := (&Lane{Name: instanceName}).Stage()
										instance = any(instanceLane)
										__gong__map_Lane[identifier] = instanceLane
									case "LaneUse":
										instanceLaneUse := (&LaneUse{Name: instanceName}).Stage()
										instance = any(instanceLaneUse)
										__gong__map_LaneUse[identifier] = instanceLaneUse
									case "Milestone":
										instanceMilestone := (&Milestone{Name: instanceName}).Stage()
										instance = any(instanceMilestone)
										__gong__map_Milestone[identifier] = instanceMilestone
									}
									__gong__map_Indentifiers_gongstructName[identifier] = gongstructName
									return
								}
							}
						}
					}
				}
				if sel := selectorExpr.Sel; sel != nil {
					// astCoordinate := astCoordinate + "\tSel" + "." + sel.Name
					// log.Println(astCoordinate)
				}
				for iteration, arg := range callExpr.Args {
					// astCoordinate := astCoordinate + "\tArg"
					switch arg := arg.(type) {
					case *ast.BasicLit:
						basicLit := arg
						// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
						// log.Println(astCoordinate)

						// first iteration should be ignored
						if iteration == 0 {
							continue
						}

						// remove first and last char
						date := basicLit.Value[1 : len(basicLit.Value)-1]
						_ = date

						var ok bool
						gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
						if !ok {
							log.Fatalln("gongstructName not found for identifier", identifier)
						}
						switch gongstructName {
						// insertion point for basic lit assignments
						case "Arrow":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Bar":
							switch fieldName {
							// insertion point for date assign code
							case "Start":
								__gong__map_Bar[identifier].Start, _ = time.Parse(
									"2006-01-02 15:04:05.999999999 -0700 MST",
									date)
							case "End":
								__gong__map_Bar[identifier].End, _ = time.Parse(
									"2006-01-02 15:04:05.999999999 -0700 MST",
									date)
							}
						case "Gantt":
							switch fieldName {
							// insertion point for date assign code
							case "ComputedStart":
								__gong__map_Gantt[identifier].ComputedStart, _ = time.Parse(
									"2006-01-02 15:04:05.999999999 -0700 MST",
									date)
							case "ComputedEnd":
								__gong__map_Gantt[identifier].ComputedEnd, _ = time.Parse(
									"2006-01-02 15:04:05.999999999 -0700 MST",
									date)
							case "ManualStart":
								__gong__map_Gantt[identifier].ManualStart, _ = time.Parse(
									"2006-01-02 15:04:05.999999999 -0700 MST",
									date)
							case "ManualEnd":
								__gong__map_Gantt[identifier].ManualEnd, _ = time.Parse(
									"2006-01-02 15:04:05.999999999 -0700 MST",
									date)
							}
						case "Group":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Lane":
							switch fieldName {
							// insertion point for date assign code
							}
						case "LaneUse":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Milestone":
							switch fieldName {
							// insertion point for date assign code
							case "Date":
								__gong__map_Milestone[identifier].Date, _ = time.Parse(
									"2006-01-02 15:04:05.999999999 -0700 MST",
									date)
							}
						}
					}
				}
			case *ast.Ident:
				// append function
				ident := fun
				_ = ident
				// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			for _, arg := range callExpr.Args {
				// astCoordinate := astCoordinate + "\tArg"
				switch arg := arg.(type) {
				case *ast.Ident:
					ident := arg
					_ = ident
					// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
					// log.Println(astCoordinate)
					var ok bool
					gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
					if !ok {
						log.Fatalln("gongstructName not found for identifier", identifier)
					}
					switch gongstructName {
					// insertion point for slice of pointers assignments
					case "Arrow":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Bar":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Gantt":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Lanes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Lane[targetIdentifier]
							__gong__map_Gantt[identifier].Lanes =
								append(__gong__map_Gantt[identifier].Lanes, target)
						case "Milestones":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Milestone[targetIdentifier]
							__gong__map_Gantt[identifier].Milestones =
								append(__gong__map_Gantt[identifier].Milestones, target)
						case "Groups":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Group[targetIdentifier]
							__gong__map_Gantt[identifier].Groups =
								append(__gong__map_Gantt[identifier].Groups, target)
						case "Arrows":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Arrow[targetIdentifier]
							__gong__map_Gantt[identifier].Arrows =
								append(__gong__map_Gantt[identifier].Arrows, target)
						}
					case "Group":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "GroupLanes":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Lane[targetIdentifier]
							__gong__map_Group[identifier].GroupLanes =
								append(__gong__map_Group[identifier].GroupLanes, target)
						}
					case "Lane":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Bars":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Bar[targetIdentifier]
							__gong__map_Lane[identifier].Bars =
								append(__gong__map_Lane[identifier].Bars, target)
						}
					case "LaneUse":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Milestone":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "LanesToDisplayMilestoneUse":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_LaneUse[targetIdentifier]
							__gong__map_Milestone[identifier].LanesToDisplayMilestoneUse =
								append(__gong__map_Milestone[identifier].LanesToDisplayMilestoneUse, target)
						}
					}
				case *ast.SelectorExpr:
					slcExpr := arg
					// astCoordinate := astCoordinate + "\tSelectorExpr"
					switch X := slcExpr.X.(type) {
					case *ast.Ident:
						ident := X
						_ = ident
						// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
						// log.Println(astCoordinate)

					}
					if Sel := slcExpr.Sel; Sel != nil {
						// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
						// log.Println(astCoordinate)
					}
				}
			}
		case *ast.BasicLit:
			// assignment to string field
			basicLit := expr
			// astCoordinate := astCoordinate + "\tBasicLit" + "." + basicLit.Value
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}

			switch gongstructName {
			// insertion point for basic lit assignments
			case "Arrow":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Arrow[identifier].Name = fielValue
				case "OptionnalColor":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Arrow[identifier].OptionnalColor = fielValue
				case "OptionnalStroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Arrow[identifier].OptionnalStroke = fielValue
				}
			case "Bar":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bar[identifier].Name = fielValue
				case "OptionnalColor":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bar[identifier].OptionnalColor = fielValue
				case "OptionnalStroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bar[identifier].OptionnalStroke = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bar[identifier].FillOpacity = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Bar[identifier].StrokeWidth = fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Bar[identifier].StrokeDashArray = fielValue
				}
			case "Gantt":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Gantt[identifier].Name = fielValue
				case "LaneHeight":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gantt[identifier].LaneHeight = fielValue
				case "RatioBarToLaneHeight":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gantt[identifier].RatioBarToLaneHeight = fielValue
				case "YTopMargin":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gantt[identifier].YTopMargin = fielValue
				case "XLeftText":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gantt[identifier].XLeftText = fielValue
				case "TextHeight":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gantt[identifier].TextHeight = fielValue
				case "XLeftLanes":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gantt[identifier].XLeftLanes = fielValue
				case "XRightMargin":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gantt[identifier].XRightMargin = fielValue
				case "ArrowLengthToTheRightOfStartBar":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gantt[identifier].ArrowLengthToTheRightOfStartBar = fielValue
				case "ArrowTipLenght":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gantt[identifier].ArrowTipLenght = fielValue
				case "TimeLine_Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Gantt[identifier].TimeLine_Color = fielValue
				case "TimeLine_FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gantt[identifier].TimeLine_FillOpacity = fielValue
				case "TimeLine_Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Gantt[identifier].TimeLine_Stroke = fielValue
				case "TimeLine_StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gantt[identifier].TimeLine_StrokeWidth = fielValue
				case "Group_Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Gantt[identifier].Group_Stroke = fielValue
				case "Group_StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gantt[identifier].Group_StrokeWidth = fielValue
				case "Group_StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Gantt[identifier].Group_StrokeDashArray = fielValue
				case "DateYOffset":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gantt[identifier].DateYOffset = fielValue
				}
			case "Group":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Group[identifier].Name = fielValue
				}
			case "Lane":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Lane[identifier].Name = fielValue
				case "Order":
					// convert string to int
					fielValue, err := strconv.ParseInt(basicLit.Value, 10, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Lane[identifier].Order = int(fielValue)
				}
			case "LaneUse":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_LaneUse[identifier].Name = fielValue
				}
			case "Milestone":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Milestone[identifier].Name = fielValue
				}
			}
		case *ast.Ident:
			// assignment to boolean field ?
			ident := expr
			_ = ident
			// astCoordinate := astCoordinate + "\tIdent" + "." + ident.Name
			// log.Println(astCoordinate)
			var ok bool
			gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
			if !ok {
				log.Fatalln("gongstructName not found for identifier", identifier)
			}
			switch gongstructName {
			// insertion point for bool & pointers assignments
			case "Arrow":
				switch fieldName {
				// insertion point for field dependant code
				case "From":
					targetIdentifier := ident.Name
					__gong__map_Arrow[identifier].From = __gong__map_Bar[targetIdentifier]
				case "To":
					targetIdentifier := ident.Name
					__gong__map_Arrow[identifier].To = __gong__map_Bar[targetIdentifier]
				}
			case "Bar":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Gantt":
				switch fieldName {
				// insertion point for field dependant code
				case "UseManualStartAndEndDates":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gantt[identifier].UseManualStartAndEndDates = fielValue
				case "AlignOnStartEndOnYearStart":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Gantt[identifier].AlignOnStartEndOnYearStart = fielValue
				}
			case "Group":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Lane":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "LaneUse":
				switch fieldName {
				// insertion point for field dependant code
				case "Lane":
					targetIdentifier := ident.Name
					__gong__map_LaneUse[identifier].Lane = __gong__map_Lane[targetIdentifier]
				}
			case "Milestone":
				switch fieldName {
				// insertion point for field dependant code
				case "DisplayVerticalBar":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Milestone[identifier].DisplayVerticalBar = fielValue
				}
			}
		case *ast.SelectorExpr:
			// assignment to enum field
			selectorExpr := expr
			// astCoordinate := astCoordinate + "\tSelectorExpr"
			switch X := selectorExpr.X.(type) {
			case *ast.Ident:
				ident := X
				_ = ident
				// astCoordinate := astCoordinate + "\tX" + "." + ident.Name
				// log.Println(astCoordinate)
			}
			if Sel := selectorExpr.Sel; Sel != nil {
				// astCoordinate := astCoordinate + "\tSel" + "." + Sel.Name
				// log.Println(astCoordinate)

				// enum field
				var ok bool
				gongstructName, ok = __gong__map_Indentifiers_gongstructName[identifier]
				if !ok {
					log.Fatalln("gongstructName not found for identifier", identifier)
				}

				// remove first and last char
				enumValue := Sel.Name
				_ = enumValue
				switch gongstructName {
				// insertion point for enums assignments
				case "Arrow":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Bar":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Gantt":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Group":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Lane":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "LaneUse":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Milestone":
					switch fieldName {
					// insertion point for enum assign code
					}
				}
			}
		}
	}
	return
}
