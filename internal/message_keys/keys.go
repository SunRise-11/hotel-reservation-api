package message_keys

// hold localization files key here
const (
	crudMessages = "CrudMessages."
	errors       = "Errors."
	users        = "Users."
	hotels       = "Hotels."
	rooms        = "Rooms."

	Created = crudMessages + "Created"
	Updated = crudMessages + "Updated"
	Deleted = crudMessages + "Deleted"

	NotFound            = errors + "NotFound"
	InternalServerError = errors + "InternalServerError"
	BadRequest          = errors + "BadRequest"

	UsernameDuplicated = users + "DuplicatedUsername"

	TypeHashotel    = "TypeHashotel"
	GradeHashotel   = "GradeHashotel"
	ValidationError = "ValidationError"
	GenderInvalid   = "GenderInvalid"

	HotelRepeatPostalCode = hotels + "RepeatPostalCode"

	InvalidRoomCleanStatus    = rooms + "InvalidCleanStatus"
	RoomTypeHasRoomErr        = rooms + "RoomTypeHasRoomErr"
	RoomHasReservationRequest = rooms + "HasReservationRequest"
)
