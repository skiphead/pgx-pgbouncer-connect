package storage

import (
	"fmt"
	"strconv"
	"strings"
)

// SelectAllColumns - Return string "SELECT * FROM <table>".
func (o *ORM) SelectAllColumns() string {

	return fmt.Sprintf("SELECT %s FROM %s", strings.Join(o.Columns, ","), o.Table)
}

// SelectWhereParam - Return string "SELECT * FROM <table> WHERE <column=param>".
func (o *ORM) SelectWhereParam(param string) string {

	return fmt.Sprintf("SELECT %s FROM %s WHERE %s=$1", strings.Join(o.Columns, ","), o.Table, param)
}

// Insert - Return string "INSERT INTO <table> (column1, column2 ...) VALUES ($1, $2 ...)"
func (o *ORM) Insert() string {
	var values []string
	if len(o.Columns) > 0 {
		for i := 0; i < len(o.Columns); i++ {
			values = append(values, "$"+strconv.Itoa(i+1))
		}
	}

	sql := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)",
		o.Table, strings.Join(o.Columns, ","), strings.Join(values, ","))

	return sql
}

func (o *ORM) setUpdate() string {
	var values []string
	if len(o.Columns) > 0 {
		for index, column := range o.Columns {
			values = append(values, fmt.Sprintf("%s=$%s", column, strconv.Itoa(index+2)))
		}
	}

	return strings.Join(values, ",")
}

// Update - Return string "ON CONFLICT <id key> DO UPDATE SET <column=$1>, <column=$2> ..."
func (o *ORM) Update() string {

	return fmt.Sprintf("UPDATE %s SET %s WHERE %s", o.Table, o.setUpdate(), o.KeyField)
}

// OnConflictDoUpdate - Return string "ON CONFLICT <id key> DO UPDATE SET <column=$1>, <column=$2> ..."
func (o *ORM) OnConflictDoUpdate() string {

	return fmt.Sprintf("ON CONFLICT (%s) DO UPDATE SET %s", o.KeyField, o.setUpdate())
}

// Delete - Return string DELETE FROM <table> WHERE <id>
func (o *ORM) Delete() string {

	return fmt.Sprintf("DELETE FROM %s WHERE %s", o.Table, o.KeyField)
}
