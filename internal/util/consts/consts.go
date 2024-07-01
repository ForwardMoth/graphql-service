package consts

const (
	BadRequest    = `Bad Request`
	InternalError = `Internal error`
	NotFound      = `Not found`
)

const (
	EmptyAuthorError         = `Author can't be empty'`
	TooMuchLengthAuthorError = `Author name can be bigger 64 symbols`
)

const (
	EmptyTextError            = `Text can't be empty'`
	TooMuchLengthTextError    = `Text can't' be bigger 3000 symbols`
	TooMuchLengthCommentError = `Comment can't' be bigger 2000 symbols`
)

const WrongIdError = `Wrong id`

const (
	CreatingPostError = `Error of creating post`
	PostNotFountError = `Posts isn't created'`
	GettingPostError  = `Posts isn't found'`
)

const (
	WrongLimitError  = `Wrong limit`
	WrongOffsetError = `Wrong offset`
)

const (
	CommentsNotAllowedError = `comments is not allowed`
	CreatingCommentError    = `error of creating comment`
)
