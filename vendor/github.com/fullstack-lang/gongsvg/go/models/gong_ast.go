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
var __gong__map_Animate = make(map[string]*Animate)
var __gong__map_Circle = make(map[string]*Circle)
var __gong__map_Ellipse = make(map[string]*Ellipse)
var __gong__map_Line = make(map[string]*Line)
var __gong__map_Path = make(map[string]*Path)
var __gong__map_Polygone = make(map[string]*Polygone)
var __gong__map_Polyline = make(map[string]*Polyline)
var __gong__map_Rect = make(map[string]*Rect)
var __gong__map_SVG = make(map[string]*SVG)
var __gong__map_Text = make(map[string]*Text)

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
									case "Animate":
										instanceAnimate := (&Animate{Name: instanceName}).Stage()
										instance = any(instanceAnimate)
										__gong__map_Animate[identifier] = instanceAnimate
									case "Circle":
										instanceCircle := (&Circle{Name: instanceName}).Stage()
										instance = any(instanceCircle)
										__gong__map_Circle[identifier] = instanceCircle
									case "Ellipse":
										instanceEllipse := (&Ellipse{Name: instanceName}).Stage()
										instance = any(instanceEllipse)
										__gong__map_Ellipse[identifier] = instanceEllipse
									case "Line":
										instanceLine := (&Line{Name: instanceName}).Stage()
										instance = any(instanceLine)
										__gong__map_Line[identifier] = instanceLine
									case "Path":
										instancePath := (&Path{Name: instanceName}).Stage()
										instance = any(instancePath)
										__gong__map_Path[identifier] = instancePath
									case "Polygone":
										instancePolygone := (&Polygone{Name: instanceName}).Stage()
										instance = any(instancePolygone)
										__gong__map_Polygone[identifier] = instancePolygone
									case "Polyline":
										instancePolyline := (&Polyline{Name: instanceName}).Stage()
										instance = any(instancePolyline)
										__gong__map_Polyline[identifier] = instancePolyline
									case "Rect":
										instanceRect := (&Rect{Name: instanceName}).Stage()
										instance = any(instanceRect)
										__gong__map_Rect[identifier] = instanceRect
									case "SVG":
										instanceSVG := (&SVG{Name: instanceName}).Stage()
										instance = any(instanceSVG)
										__gong__map_SVG[identifier] = instanceSVG
									case "Text":
										instanceText := (&Text{Name: instanceName}).Stage()
										instance = any(instanceText)
										__gong__map_Text[identifier] = instanceText
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
						case "Animate":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Circle":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Ellipse":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Line":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Path":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Polygone":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Polyline":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Rect":
							switch fieldName {
							// insertion point for date assign code
							}
						case "SVG":
							switch fieldName {
							// insertion point for date assign code
							}
						case "Text":
							switch fieldName {
							// insertion point for date assign code
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
					case "Animate":
						switch fieldName {
						// insertion point for slice of pointers assign code
						}
					case "Circle":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animations":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_Circle[identifier].Animations =
								append(__gong__map_Circle[identifier].Animations, target)
						}
					case "Ellipse":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_Ellipse[identifier].Animates =
								append(__gong__map_Ellipse[identifier].Animates, target)
						}
					case "Line":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_Line[identifier].Animates =
								append(__gong__map_Line[identifier].Animates, target)
						}
					case "Path":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_Path[identifier].Animates =
								append(__gong__map_Path[identifier].Animates, target)
						}
					case "Polygone":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_Polygone[identifier].Animates =
								append(__gong__map_Polygone[identifier].Animates, target)
						}
					case "Polyline":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_Polyline[identifier].Animates =
								append(__gong__map_Polyline[identifier].Animates, target)
						}
					case "Rect":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animations":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_Rect[identifier].Animations =
								append(__gong__map_Rect[identifier].Animations, target)
						}
					case "SVG":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Rects":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Rect[targetIdentifier]
							__gong__map_SVG[identifier].Rects =
								append(__gong__map_SVG[identifier].Rects, target)
						case "Texts":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Text[targetIdentifier]
							__gong__map_SVG[identifier].Texts =
								append(__gong__map_SVG[identifier].Texts, target)
						case "Circles":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Circle[targetIdentifier]
							__gong__map_SVG[identifier].Circles =
								append(__gong__map_SVG[identifier].Circles, target)
						case "Lines":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Line[targetIdentifier]
							__gong__map_SVG[identifier].Lines =
								append(__gong__map_SVG[identifier].Lines, target)
						case "Ellipses":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Ellipse[targetIdentifier]
							__gong__map_SVG[identifier].Ellipses =
								append(__gong__map_SVG[identifier].Ellipses, target)
						case "Polylines":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Polyline[targetIdentifier]
							__gong__map_SVG[identifier].Polylines =
								append(__gong__map_SVG[identifier].Polylines, target)
						case "Polygones":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Polygone[targetIdentifier]
							__gong__map_SVG[identifier].Polygones =
								append(__gong__map_SVG[identifier].Polygones, target)
						case "Paths":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Path[targetIdentifier]
							__gong__map_SVG[identifier].Paths =
								append(__gong__map_SVG[identifier].Paths, target)
						}
					case "Text":
						switch fieldName {
						// insertion point for slice of pointers assign code
						case "Animates":
							// remove first and last char
							targetIdentifier := ident.Name
							target := __gong__map_Animate[targetIdentifier]
							__gong__map_Text[identifier].Animates =
								append(__gong__map_Text[identifier].Animates, target)
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
			case "Animate":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Animate[identifier].Name = fielValue
				case "AttributeName":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Animate[identifier].AttributeName = fielValue
				case "Values":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Animate[identifier].Values = fielValue
				case "Dur":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Animate[identifier].Dur = fielValue
				case "RepeatCount":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Animate[identifier].RepeatCount = fielValue
				}
			case "Circle":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].Name = fielValue
				case "CX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].CX = fielValue
				case "CY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].CY = fielValue
				case "Radius":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].Radius = fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].FillOpacity = fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Circle[identifier].StrokeWidth = fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].StrokeDashArray = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Circle[identifier].Transform = fielValue
				}
			case "Ellipse":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Ellipse[identifier].Name = fielValue
				case "CX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Ellipse[identifier].CX = fielValue
				case "CY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Ellipse[identifier].CY = fielValue
				case "RX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Ellipse[identifier].RX = fielValue
				case "RY":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Ellipse[identifier].RY = fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Ellipse[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Ellipse[identifier].FillOpacity = fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Ellipse[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Ellipse[identifier].StrokeWidth = fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Ellipse[identifier].StrokeDashArray = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Ellipse[identifier].Transform = fielValue
				}
			case "Line":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Line[identifier].Name = fielValue
				case "X1":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Line[identifier].X1 = fielValue
				case "Y1":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Line[identifier].Y1 = fielValue
				case "X2":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Line[identifier].X2 = fielValue
				case "Y2":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Line[identifier].Y2 = fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Line[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Line[identifier].FillOpacity = fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Line[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Line[identifier].StrokeWidth = fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Line[identifier].StrokeDashArray = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Line[identifier].Transform = fielValue
				}
			case "Path":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Path[identifier].Name = fielValue
				case "Definition":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Path[identifier].Definition = fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Path[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Path[identifier].FillOpacity = fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Path[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Path[identifier].StrokeWidth = fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Path[identifier].StrokeDashArray = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Path[identifier].Transform = fielValue
				}
			case "Polygone":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polygone[identifier].Name = fielValue
				case "Points":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polygone[identifier].Points = fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polygone[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Polygone[identifier].FillOpacity = fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polygone[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Polygone[identifier].StrokeWidth = fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polygone[identifier].StrokeDashArray = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polygone[identifier].Transform = fielValue
				}
			case "Polyline":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polyline[identifier].Name = fielValue
				case "Points":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polyline[identifier].Points = fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polyline[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Polyline[identifier].FillOpacity = fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polyline[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Polyline[identifier].StrokeWidth = fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polyline[identifier].StrokeDashArray = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Polyline[identifier].Transform = fielValue
				}
			case "Rect":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rect[identifier].Name = fielValue
				case "X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].X = fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].Y = fielValue
				case "Width":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].Width = fielValue
				case "Height":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].Height = fielValue
				case "RX":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].RX = fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rect[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].FillOpacity = fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rect[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Rect[identifier].StrokeWidth = fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rect[identifier].StrokeDashArray = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Rect[identifier].Transform = fielValue
				}
			case "SVG":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_SVG[identifier].Name = fielValue
				}
			case "Text":
				switch fieldName {
				// insertion point for field dependant code
				case "Name":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].Name = fielValue
				case "X":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Text[identifier].X = fielValue
				case "Y":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Text[identifier].Y = fielValue
				case "Content":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].Content = fielValue
				case "Color":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].Color = fielValue
				case "FillOpacity":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Text[identifier].FillOpacity = fielValue
				case "Stroke":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].Stroke = fielValue
				case "StrokeWidth":
					// convert string to float64
					fielValue, err := strconv.ParseFloat(basicLit.Value, 64)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_Text[identifier].StrokeWidth = fielValue
				case "StrokeDashArray":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].StrokeDashArray = fielValue
				case "Transform":
					// remove first and last char
					fielValue := basicLit.Value[1 : len(basicLit.Value)-1]
					__gong__map_Text[identifier].Transform = fielValue
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
			case "Animate":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Circle":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Ellipse":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Line":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Path":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Polygone":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Polyline":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "Rect":
				switch fieldName {
				// insertion point for field dependant code
				}
			case "SVG":
				switch fieldName {
				// insertion point for field dependant code
				case "Display":
					// convert string to boolean
					fielValue, err := strconv.ParseBool(ident.Name)
					if err != nil {
						log.Fatalln(err)
					}
					__gong__map_SVG[identifier].Display = fielValue
				}
			case "Text":
				switch fieldName {
				// insertion point for field dependant code
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
				case "Animate":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Circle":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Ellipse":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Line":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Path":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Polygone":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Polyline":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Rect":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "SVG":
					switch fieldName {
					// insertion point for enum assign code
					}
				case "Text":
					switch fieldName {
					// insertion point for enum assign code
					}
				}
			}
		}
	}
	return
}
