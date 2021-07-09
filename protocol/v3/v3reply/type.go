package v3reply

import "github.com/wwq-2020/go-redis/protocol"

// consts
const (
	ReplyTypeBlobError      protocol.ReplyType = "BlobError"
	ReplyTypeBlobString     protocol.ReplyType = "BlobString"
	ReplyTypeBoolean        protocol.ReplyType = "Boolean"
	ReplyTypeDouble         protocol.ReplyType = "Double"
	ReplyTypeMap            protocol.ReplyType = "Map"
	ReplyTypeNull           protocol.ReplyType = "NULL"
	ReplyTypeNumber         protocol.ReplyType = "Number"
	ReplyTypePush           protocol.ReplyType = "Push"
	ReplyTypeSimpleError    protocol.ReplyType = "SimpleError"
	ReplyTypeSimpleString   protocol.ReplyType = "SimpleString"
	ReplyTypeStream         protocol.ReplyType = "Stream"
	ReplyTypeVerbatimString protocol.ReplyType = "VerbatimString"
	ReplyTypeBigNumber      protocol.ReplyType = "BigNumber"
	ReplyTypeSet            protocol.ReplyType = "Set"
	ReplyArray              protocol.ReplyType = "Array"
	ReplyStatus             protocol.ReplyType = "Status"
)

// consts
const (
	typeBlobString     = '$'     // $<length>\r\n<bytes>\r\n
	typeSimpleString   = '+'     // +<string>\r\n
	typeSimpleError    = '-'     // -<string>\r\n
	typeNumber         = ':'     // :<number>\r\n
	typeNull           = '_'     // _\r\n
	typeDouble         = ','     // ,<floating-point-number>\r\n
	typeBoolean        = '#'     // #t\r\n or #f\r\n
	typeBlobError      = '!'     // !<length>\r\n<bytes>\r\n
	typeVerbatimString = '='     // =<length>\r\n<format(3 bytes):><bytes>\r\n
	typeBigNumber      = '('     // (<big number>\n
	typeArray          = '*'     // *<elements number>\r\n... numelements other types ...
	typeMap            = '%'     // %<elements number>\r\n... numelements key/value pair of other types ...
	typeSet            = '~'     // ~<elements number>\r\n... numelements other types ...
	typeAttribute      = '|'     // |~<elements number>\r\n... numelements map type ...
	typePush           = '>'     // ><elements number>\r\n<first item is String>\r\n... numelements-1 other types ...
	typeStream         = "$EOF:" // $EOF:<40 bytes marker><CR><LF>... any number of bytes of data here not containing the marker ...<40 bytes marker>
)
