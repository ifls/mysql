// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2017 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

import (
	"database/sql"
	"reflect"
)

// 获取字段类型对应的数据库类型名
func (mf *mysqlField) typeDatabaseName() string {
	switch mf.fieldType {
	// 数值
	case fieldTypeTiny: //1B
		return "TINYINT"
	case fieldTypeShort: //2B
		return "SMALLINT"
	case fieldTypeInt24: //3B
		return "MEDIUMINT"
	case fieldTypeLong: //4B
		return "INT"
	case fieldTypeLongLong: // 8B
		return "BIGINT"
	case fieldTypeFloat: //4B
		return "FLOAT"
	case fieldTypeDouble: //8B
		return "DOUBLE"
	case fieldTypeDecimal: // decimal(m, d) (max(m,d) + 2)B
		return "DECIMAL"
	// 字符串
	case fieldTypeString: // 定长 字符串 [0-255]B
		if mf.charSet == collations[binaryCollation] {
			return "BINARY" // 二进制, []byte
		}
		return "CHAR" //字符数组, 需要指定编码集
	case fieldTypeVarChar: //同下
		if mf.charSet == collations[binaryCollation] {
			return "VARBINARY"
		}
		return "VARCHAR"
	case fieldTypeVarString: //变长字符串 [0-65535]B char(n) varchar(n) 表示 字符的个数, 不是字节数
		if mf.charSet == collations[binaryCollation] {
			return "VARBINARY"
		}
		return "VARCHAR"
	case fieldTypeTinyBLOB: // 定长字节块 [0-255]B
		if mf.charSet != collations[binaryCollation] {
			return "TINYTEXT"
		}
		return "TINYBLOB"
	case fieldTypeBLOB: //[0-2^16]B
		if mf.charSet != collations[binaryCollation] {
			return "TEXT" //text 是文本
		}
		return "BLOB" // blob 是二进制块
	case fieldTypeMediumBLOB: // [0-2^24]B
		if mf.charSet != collations[binaryCollation] {
			return "MEDIUMTEXT"
		}
		return "MEDIUMBLOB"
	case fieldTypeLongBLOB: // [0-2^32]B
		if mf.charSet != collations[binaryCollation] {
			return "LONGTEXT"
		}
		return "LONGBLOB"

	// 日期
	case fieldTypeYear: //2B
		return "YEAR"
	case fieldTypeDate: //3B YYYY-MM-DD
		return "DATE"
	case fieldTypeTime: //3B HH:MM:SS
		return "TIME"
	case fieldTypeTimestamp: //4B unix 时间戳 秒级
		return "TIMESTAMP"
	case fieldTypeDateTime: //8B YYYY-MM-DD HH:MM:SS
		return "DATETIME"
	// 特殊
	case fieldTypeBit: // 特殊的 BIT(m) 存储 m 位, 提供位操作
		return "BIT"
	case fieldTypeEnum: // 特殊 字符串类型的枚举定义
		return "ENUM"
	case fieldTypeGeometry: // 地理位置  point
		return "GEOMETRY"
	case fieldTypeJSON: //json 不能设定长度
		return "JSON"
	case fieldTypeNewDate: //?
		return "DATE"
	case fieldTypeNewDecimal: //?
		return "DECIMAL"
	case fieldTypeSet: // 集合
		return "SET"

	// NULL
	case fieldTypeNULL: //NULL
		return "NULL"
	default:
		return ""
	}
}

var (
	// 反射类型
	scanTypeFloat32   = reflect.TypeOf(float32(0))
	scanTypeFloat64   = reflect.TypeOf(float64(0))
	scanTypeInt8      = reflect.TypeOf(int8(0))
	scanTypeInt16     = reflect.TypeOf(int16(0))
	scanTypeInt32     = reflect.TypeOf(int32(0))
	scanTypeInt64     = reflect.TypeOf(int64(0))
	scanTypeNullFloat = reflect.TypeOf(sql.NullFloat64{})
	scanTypeNullInt   = reflect.TypeOf(sql.NullInt64{})
	scanTypeNullTime  = reflect.TypeOf(NullTime{})
	scanTypeUint8     = reflect.TypeOf(uint8(0))
	scanTypeUint16    = reflect.TypeOf(uint16(0))
	scanTypeUint32    = reflect.TypeOf(uint32(0))
	scanTypeUint64    = reflect.TypeOf(uint64(0))
	scanTypeRawBytes  = reflect.TypeOf(sql.RawBytes{})
	scanTypeUnknown   = reflect.TypeOf(new(interface{}))
)

type mysqlField struct {
	tableName string    //表名
	name      string    //字段名
	length    uint32    //数据长度 例如 varchar(n)
	flags     fieldFlag //字段标志, 主键, 飞空
	fieldType fieldType //字段类型
	decimals  byte      //小数点后的位数
	charSet   uint8     //字符集
}

// 返回反射类型
func (mf *mysqlField) scanType() reflect.Type {
	switch mf.fieldType {
	case fieldTypeTiny:
		if mf.flags&flagNotNULL != 0 {
			if mf.flags&flagUnsigned != 0 {
				return scanTypeUint8
			}
			return scanTypeInt8
		}
		return scanTypeNullInt

	case fieldTypeShort, fieldTypeYear:
		if mf.flags&flagNotNULL != 0 {
			if mf.flags&flagUnsigned != 0 {
				return scanTypeUint16
			}
			return scanTypeInt16
		}
		return scanTypeNullInt

	case fieldTypeInt24, fieldTypeLong:
		if mf.flags&flagNotNULL != 0 {
			if mf.flags&flagUnsigned != 0 {
				return scanTypeUint32
			}
			return scanTypeInt32
		}
		return scanTypeNullInt

	case fieldTypeLongLong:
		if mf.flags&flagNotNULL != 0 {
			if mf.flags&flagUnsigned != 0 {
				return scanTypeUint64
			}
			return scanTypeInt64
		}
		return scanTypeNullInt

	case fieldTypeFloat:
		if mf.flags&flagNotNULL != 0 {
			return scanTypeFloat32
		}
		return scanTypeNullFloat

	case fieldTypeDouble:
		if mf.flags&flagNotNULL != 0 {
			return scanTypeFloat64
		}
		return scanTypeNullFloat

	case fieldTypeDecimal, fieldTypeNewDecimal, fieldTypeVarChar,
		fieldTypeBit, fieldTypeEnum, fieldTypeSet, fieldTypeTinyBLOB,
		fieldTypeMediumBLOB, fieldTypeLongBLOB, fieldTypeBLOB,
		fieldTypeVarString, fieldTypeString, fieldTypeGeometry, fieldTypeJSON,
		fieldTypeTime:
		return scanTypeRawBytes

	case fieldTypeDate, fieldTypeNewDate,
		fieldTypeTimestamp, fieldTypeDateTime:
		// NullTime is always returned for more consistent behavior as it can
		// handle both cases of parseTime regardless if the field is nullable.
		return scanTypeNullTime

	default:
		return scanTypeUnknown
	}
}
