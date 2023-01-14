package models

import (
	"errors"
	"go/ast"
	"go/doc/comment"
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
func ParseAstFile(pathToFile string) error {

	// map to store renaming docLink
	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	Stage.Map_DocLink_Renaming = make(map[string]string, 0)

	fileOfInterest, err := filepath.Abs(pathToFile)
	if err != nil {
		return errors.New("Path does not exist %s ;" + fileOfInterest)
	}

	fset := token.NewFileSet()
	startParser := time.Now()
	inFile, errParser := parser.ParseFile(fset, fileOfInterest, nil, parser.ParseComments)
	log.Printf("Parser took %s", time.Since(startParser))

	if errParser != nil {
		return errors.New("Unable to parser " + errParser.Error())
	}

	// if there is a meta package import, it is the third import
	if len(inFile.Imports) > 3 {
		log.Fatalln("Too many imports in file", fileOfInterest)
	}
	stage := &Stage
	_ = stage
	if len(inFile.Imports) == 3 {
		Stage.MetaPackageImportAlias = inFile.Imports[2].Name.Name
		Stage.MetaPackageImportPath = inFile.Imports[2].Path.Value
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
						// Create an ast.CommentMap from the ast.File's comments.
						// This helps keeping the association between comments
						// and AST nodes.
						cmap := ast.NewCommentMap(fset, inFile, inFile.Comments)
						astCoordinate := "\tAssignStmt: "
						// log.Println(// astCoordinate)
						assignStmt := stmt
						instance, id, gongstruct, fieldName :=
							UnmarshallGongstructStaging(
								&cmap, assignStmt, astCoordinate)
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
				case *ast.ValueSpec:
					ident := spec.Names[0]
					_ = ident
					if !strings.HasPrefix(ident.Name, "map_DocLink_Identifier") {
						continue
					}
					switch compLit := spec.Values[0].(type) {
					case *ast.CompositeLit:
						var key string
						_ = key
						var value string
						_ = value
						for _, elt := range compLit.Elts {

							// each elt is an expression for struct or for field such as
							// for struct
							//
							//         "dummy.Dummy": &(dummy.Dummy{})
							//
							// or, for field
							//
							//          "dummy.Dummy.Name": (dummy.Dummy{}).Name,
							//
							// first node in the AST is a key value expression
							var ok bool
							var kve *ast.KeyValueExpr
							if kve, ok = elt.(*ast.KeyValueExpr); !ok {
								log.Fatal("Expression should be key value expression" +
									fset.Position(kve.Pos()).String())
							}

							switch bl := kve.Key.(type) {
							case *ast.BasicLit:
								key = bl.Value // "\"dumm.Dummy\"" is the value

								// one remove the ambracing double quotes
								key = strings.TrimPrefix(key, "\"")
								key = strings.TrimSuffix(key, "\"")
							}

							var isFieldEntry bool
							var fieldName string
							var ue *ast.UnaryExpr
							if ue, ok = kve.Value.(*ast.UnaryExpr); !ok {
								isFieldEntry = true
							}

							var se2 *ast.SelectorExpr
							if isFieldEntry {
								if se2, ok = kve.Value.(*ast.SelectorExpr); !ok {
									log.Fatal("Expression should be a selector expression" +
										fset.Position(kve.Pos()).String())
								}
								fieldName = se2.Sel.Name
							}

							var pe *ast.ParenExpr
							if !isFieldEntry {
								if pe, ok = ue.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							} else {
								if pe, ok = se2.X.(*ast.ParenExpr); !ok {
									log.Fatal("Expression should be parenthese expression" +
										fset.Position(ue.Pos()).String())
								}
							}

							// expect a Composite Litteral with no Element <type>{}
							var cl *ast.CompositeLit
							if cl, ok = pe.X.(*ast.CompositeLit); !ok {
								log.Fatal("Expression should be a composite lit" +
									fset.Position(pe.Pos()).String())
							}

							var se *ast.SelectorExpr
							if se, ok = cl.Type.(*ast.SelectorExpr); !ok {
								log.Fatal("Expression should be a selector" +
									fset.Position(cl.Pos()).String())
							}

							var id *ast.Ident
							if id, ok = se.X.(*ast.Ident); !ok {
								log.Fatal("Expression should be an ident" +
									fset.Position(se.Pos()).String())
							}
							docLink := id.Name + "." + se.Sel.Name

							if isFieldEntry {
								docLink += "." + fieldName
							}

							// if map_DocLink_Identifier has the same ident, this means
							// that no renaming has occured since the last processing of the
							// file. But it is neccessary to keep it in memory for the
							// marshalling
							if docLink == key {
								// continue
							}

							// otherwise, one stores the new ident (after renaming) in the
							// renaming map
							Stage.Map_DocLink_Renaming[key] = docLink
						}
					}
				}
			}
		}

	}
	return nil
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

// Parser needs to be configured for having the [Name1.Name2] or [pkg.Name1] ...
// to be recognized as a proper identifier.
// While this was introduced in go 1.19, it is not yet implemented in
// gopls (see [issue](https://github.com/golang/go/issues/57559)
func lookupPackage(name string) (importPath string, ok bool) {
	if name == Stage.MetaPackageImportAlias {
		return Stage.MetaPackageImportAlias, true
	}
	return comment.DefaultLookupPackage(name)
}
func lookupSym(recv, name string) (ok bool) {
	if recv == "" {
		return true
	}
	return false
}

// UnmarshallGoStaging unmarshall a go assign statement
func UnmarshallGongstructStaging(cmap *ast.CommentMap, assignStmt *ast.AssignStmt, astCoordinate_ string) (
	instance any,
	identifier string,
	gongstructName string,
	fieldName string) {

	// used for debug purposes
	astCoordinate := "\tAssignStmt: "

	//
	// First parse all comment groups in the assignement
	// if a comment "//gong:ident [DocLink]" is met and is followed by a string assignement.
	// modify the following AST assignement to assigns the DocLink text to the string value
	//

	// get the comment group of the assigStmt
	commentGroups := (*cmap)[assignStmt]
	// get the the prefix
	var hasGongIdentDirective bool
	var commentText string
	var docLinkText string
	for _, commentGroup := range commentGroups {
		for _, comment := range commentGroup.List {
			if strings.HasPrefix(comment.Text, "//gong:ident") {
				hasGongIdentDirective = true
				commentText = comment.Text
			}
		}
	}
	if hasGongIdentDirective {
		// parser configured to find doclinks
		var docLinkFinder comment.Parser
		docLinkFinder.LookupPackage = lookupPackage
		docLinkFinder.LookupSym = lookupSym
		doc := docLinkFinder.Parse(commentText)

		for _, block := range doc.Content {
			switch paragraph := block.(type) {
			case *comment.Paragraph:
				_ = paragraph
				for _, text := range paragraph.Text {
					switch docLink := text.(type) {
					case *comment.DocLink:
						if docLink.Recv == "" {
							docLinkText = docLink.ImportPath + "." + docLink.Name
						} else {
							docLinkText = docLink.ImportPath + "." + docLink.Recv + "." + docLink.Name
						}

						// we check wether the doc link has been renamed
						// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
						if renamed, ok := (Stage.Map_DocLink_Renaming)[docLinkText]; ok {
							docLinkText = renamed
						}
					}
				}
			}
		}
	}

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

			// substitute the RHS part of the assignment if a //gong:ident directive is met
			if hasGongIdentDirective {
				basicLit.Value = "[" + docLinkText + "]"
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
