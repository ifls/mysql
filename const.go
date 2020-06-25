// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 The Go-MySQL-Driver Authors. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

const (
	defaultAuthPlugin       = "mysql_native_password"
	defaultMaxAllowedPacket = 4 << 20 // 4 MiB
	minProtocolVersion      = 10
	maxPacketSize           = 1<<24 - 1
	timeFormat              = "2006-01-02 15:04:05.999999"
)

// MySQL constants documentation:
// http://dev.mysql.com/doc/internals/en/client-server-protocol.html

// 包头标志
const (
	iOK           byte = 0x00		//ok 包括 PREPARE_OK

	iAuthMoreData byte = 0x01		//
	iLocalInFile  byte = 0xfb

	iEOF          byte = 0xfe
	iERR          byte = 0xff
)

// https://dev.mysql.com/doc/internals/en/capability-flags.html#packet-Protocol::CapabilityFlags
// capability flags 能力标记 用于表明支持的能用的特性
type clientFlag uint32

const (
	clientLongPassword clientFlag = 1 << iota	//强密码
	clientFoundRows
	clientLongFlag
	clientConnectWithDB
	clientNoSchema
	clientCompress
	clientODBC
	clientLocalFiles
	clientIgnoreSpace
	clientProtocol41
	clientInteractive
	clientSSL
	clientIgnoreSIGPIPE
	clientTransactions
	clientReserved
	clientSecureConn
	clientMultiStatements
	clientMultiResults
	clientPSMultiResults
	clientPluginAuth
	clientConnectAttrs
	clientPluginAuthLenEncClientData
	clientCanHandleExpiredPasswords
	clientSessionTrack
	clientDeprecateEOF
)

//命令列表 https://dev.mysql.com/doc/internals/en/text-protocol.html
const (
								// COM_SLEEP 内部服务器命令
	comQuit byte = iota + 1		//关闭/退出连接
	comInitDB					//切换数据库
	comQuery					//文本形式 立刻执行 SQL 查询 包括 select 和 增删改
	comFieldList				//获取数据表字段信息
	comCreateDB					//创建数据库
	comDropDB					//删除数据库
	comRefresh					//清楚缓存
	comShutdown					//停止服务器
	comStatistics				//获取服务器统计信息
	comProcessInfo				//获取当前连接的列表
	comConnect					// (服务器内部命令)
	comProcessKill				//中断一个连接
	comDebug					//设置调试模式,保存服务器调试信息
	comPing						//测试 连通性
	comTime						// (服务器内部命令)
	comDelayedInsert			// (服务器内部命令)
	comChangeUser				// 重新登录(不断连接)
	comBinlogDump				//获取二进制日志信息
	comTableDump				//获取数据表结构信息
	comConnectOut				// (服务器内部命令)
	comRegisterSlave			// (从服务器向主服务器注册)
	comStmtPrepare				//预处理 SQL 语句
	comStmtExecute				//执行预处理语句
	comStmtSendLongData			//发送 BLOB 类型的数据
	comStmtClose				//销毁预处理语句
	comStmtReset				//清楚预处理语句参数缓存
	comSetOption				//设置语句选项
	comStmtFetch				//获取预处理语句的执行结果
)

// https://dev.mysql.com/doc/internals/en/com-query-response.html#packet-Protocol::ColumnType
type fieldType byte

const (
	fieldTypeDecimal fieldType = iota
	fieldTypeTiny
	fieldTypeShort
	fieldTypeLong
	fieldTypeFloat
	fieldTypeDouble
	fieldTypeNULL
	fieldTypeTimestamp
	fieldTypeLongLong
	fieldTypeInt24
	fieldTypeDate
	fieldTypeTime
	fieldTypeDateTime
	fieldTypeYear
	fieldTypeNewDate
	fieldTypeVarChar
	fieldTypeBit
)
const (
	fieldTypeJSON fieldType = iota + 0xf5
	fieldTypeNewDecimal
	fieldTypeEnum
	fieldTypeSet
	fieldTypeTinyBLOB
	fieldTypeMediumBLOB
	fieldTypeLongBLOB
	fieldTypeBLOB
	fieldTypeVarString
	fieldTypeString
	fieldTypeGeometry
)

type fieldFlag uint16

const (
	flagNotNULL fieldFlag = 1 << iota
	flagPriKey
	flagUniqueKey
	flagMultipleKey
	flagBLOB
	flagUnsigned
	flagZeroFill
	flagBinary
	flagEnum
	flagAutoIncrement
	flagTimestamp
	flagSet
	flagUnknown1
	flagUnknown2
	flagUnknown3
	flagUnknown4
)

// http://dev.mysql.com/doc/internals/en/status-flags.html
type statusFlag uint16

const (
	statusInTrans statusFlag = 1 << iota		//a transaction is active
	statusInAutocommit							//auto commit is enabled
	statusReserved // Not in documentation
	statusMoreResultsExists
	statusNoGoodIndexUsed
	statusNoIndexUsed
	statusCursorExists
	statusLastRowSent
	statusDbDropped
	statusNoBackslashEscapes
	statusMetadataChanged
	statusQueryWasSlow
	statusPsOutParams
	statusInTransReadonly						//in a read-only transaction
	statusSessionStateChanged					//connection state information has changed
)

const (
	cachingSha2PasswordRequestPublicKey          = 2
	cachingSha2PasswordFastAuthSuccess           = 3
	cachingSha2PasswordPerformFullAuthentication = 4
)
