package v3

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
