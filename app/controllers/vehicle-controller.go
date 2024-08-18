package controller

import (
	"framework/internal/app/db"
	"framework/internal/app/request"
	"framework/internal/app/session"
	"framework/internal/app/view"
	"strings"
	"unicode"

	builder "github.com/olbrichattila/gosqlbuilder/pkg"
)

func DisplayAllMakes(r request.Requester, db db.DBer, v view.Viewer, sqlBuilder builder.Builder, s session.Sessioner) (string, error) {
	defer db.Close()

	sqlBuilder.Select("car_make").
		Fields("make").
		OrderBy("make")

	sql, err := sqlBuilder.AsSQL()
	if err != nil {
		return "", err
	}

	report := make([]map[string]interface{}, 0)
	res := db.QueryAll(sql, sqlBuilder.GetParams()...)
	if db.GetLastError() != nil {
		return "", db.GetLastError()
	}
	for ret := range res {
		report = append(report, ret)
	}

	if db.GetLastError() != nil {
		return "", db.GetLastError()
	}

	templateFiles := []string{
		"make.html",
	}

	return v.Render(templateFiles, report), nil
}

func DisplayAllSubModels(r request.Requester, db db.DBer, v view.Viewer, sqlBuilder builder.Builder, s session.Sessioner) (string, error) {
	defer db.Close()

	sqlBuilder.Select("car_basemodel").
		Fields("make", "basemodel").
		Where("make", "=", r.GetOne("make", "")).
		OrderBy("basemodel")

	sql, err := sqlBuilder.AsSQL()
	if err != nil {
		return "", err
	}

	report := make([]map[string]interface{}, 0)
	res := db.QueryAll(sql, sqlBuilder.GetParams()...)
	if db.GetLastError() != nil {
		return "", db.GetLastError()
	}
	for ret := range res {
		report = append(report, ret)
	}

	if db.GetLastError() != nil {
		return "", db.GetLastError()
	}

	templateFiles := []string{
		"basemodel.html",
	}

	return v.Render(templateFiles, report), nil
}

func DisplayAllModels(r request.Requester, db db.DBer, v view.Viewer, sqlBuilder builder.Builder, s session.Sessioner) (string, error) {
	defer db.Close()

	sqlBuilder.Select("car_model").
		Fields("make", "basemodel", "model").
		Where("basemodel", "=", r.GetOne("basemodel", "")).
		Where("make", "=", r.GetOne("make", "")).
		OrderBy("model")

	sql, err := sqlBuilder.AsSQL()
	if err != nil {
		return "", err
	}

	report := make([]map[string]interface{}, 0)
	res := db.QueryAll(sql, sqlBuilder.GetParams()...)
	if db.GetLastError() != nil {
		return "", db.GetLastError()
	}
	for ret := range res {
		report = append(report, ret)
	}

	if db.GetLastError() != nil {
		return "", db.GetLastError()
	}

	templateFiles := []string{
		"model.html",
	}

	return v.Render(templateFiles, report), nil
}

func DisplayAllFuelType(r request.Requester, db db.DBer, v view.Viewer, sqlBuilder builder.Builder, s session.Sessioner) (string, error) {
	defer db.Close()

	sqlBuilder.Select("car_fuel_type").
		Fields("make", "basemodel", "model", "fuel_type").
		Where("make", "=", r.GetOne("make", "")).
		Where("basemodel", "=", r.GetOne("basemodel", "")).
		Where("model", "=", r.GetOne("model", "")).
		OrderBy("fuel_type")

	sql, err := sqlBuilder.AsSQL()
	if err != nil {
		return "", err
	}

	report := make([]map[string]interface{}, 0)
	res := db.QueryAll(sql, sqlBuilder.GetParams()...)
	if db.GetLastError() != nil {
		return "", db.GetLastError()
	}
	for ret := range res {
		report = append(report, ret)
	}

	if db.GetLastError() != nil {
		return "", db.GetLastError()
	}

	templateFiles := []string{
		"fuel_type.html",
	}

	return v.Render(templateFiles, report), nil
}

func DisplayAllYear(r request.Requester, db db.DBer, v view.Viewer, sqlBuilder builder.Builder, s session.Sessioner) (string, error) {
	defer db.Close()

	sqlBuilder.Select("car_year").
		Fields("make", "basemodel", "model", "fuel_type", "year").
		Where("make", "=", r.GetOne("make", "")).
		Where("basemodel", "=", r.GetOne("basemodel", "")).
		Where("model", "=", r.GetOne("model", "")).
		Where("fuel_type", "=", r.GetOne("fuel_type", "")).
		OrderBy("fuel_type")

	sql, err := sqlBuilder.AsSQL()
	if err != nil {
		return "", err
	}

	report := make([]map[string]interface{}, 0)
	res := db.QueryAll(sql, sqlBuilder.GetParams()...)
	if db.GetLastError() != nil {
		return "", db.GetLastError()
	}
	for ret := range res {
		report = append(report, ret)
	}

	if db.GetLastError() != nil {
		return "", db.GetLastError()
	}

	templateFiles := []string{
		"year.html",
	}

	return v.Render(templateFiles, report), nil
}

func DisplayVehicles(r request.Requester, db db.DBer, v view.Viewer, sqlBuilder builder.Builder, s session.Sessioner) (string, error) {
	defer db.Close()

	sqlBuilder.Select("vehicles").
		Where("make", "=", r.GetOne("make", "")).
		Where("basemodel", "=", r.GetOne("basemodel", "")).
		Where("model", "=", r.GetOne("model", "")).
		Where("fuel_type", "=", r.GetOne("fuel_type", "")).
		Where("year", "=", r.GetOne("year", ""))

	sql, err := sqlBuilder.AsSQL()
	if err != nil {
		return "", err
	}

	report := make([]map[string]interface{}, 0)
	res := db.QueryAll(sql, sqlBuilder.GetParams()...)
	if db.GetLastError() != nil {
		return "", db.GetLastError()
	}

	row := make(map[string]interface{}, 0)
	for ret := range res {
		for fieldName, fieldValue := range ret {
			r := []rune(strings.ReplaceAll(fieldName, "_", " "))
			r[0] = unicode.ToUpper(r[0])

			row[string(r)] = fieldValue
		}
		report = append(report, row)
	}

	if db.GetLastError() != nil {
		return "", db.GetLastError()
	}

	templateFiles := []string{
		"vehicles.html",
	}

	return v.Render(templateFiles, report), nil
}
