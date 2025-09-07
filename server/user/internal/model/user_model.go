package model

type RegisterRequest struct {
	NamaLengkap        string `json:"nama_lengkap" validate:"required,max=100"`
	Username           string `json:"username" validate:"required,max=50"`
	Email              string `json:"email" validate:"required,max=100,email"`
	NoHP               string `json:"no_hp" validate:"required,max=12"`
	Password           string `json:"password" validate:"passwd,required,min=5"`
	KonfirmasiPassword string `json:"konfirmasi_password" validate:"eqfield=password"`
}

type DataResponse struct {
	Data *UserResponse `json:"data"`
}

type UpdateAlamatRequest struct {
	NamaJalan string `json:"nama_jalan" validate:"required,max=100"`
	RT        int    `json:"rt" validate:"required"`
	RW        int    `json:"rw" validate:"required"`
	Kelurahan string `json:"kelurahan" validate:"required,max=100"`
	Kecamatan string `json:"kecamatan" validate:"required,max=100"`
	Kota      string `json:"kota" validate:"required,max=100"`
	KodePos   int    `json:"kode_pos" validate:"required"`
}

type UpdateUserRequest struct {
	NamaLengkap        string               `json:"nama_lengkap" validate:"required,max=100"`
	Username           string               `json:"username" validate:"required,max=50"`
	Email              string               `json:"email" validate:"required,max=100,email"`
	Alamat             *UpdateAlamatRequest `json:"alamat"`
	NoHP               string               `json:"no_hp" validate:"required,max=12"`
	Password           string               `json:"password" validate:"passwd,required,min=5"`
	KonfirmasiPassword string               `json:"konfirmasi_password" validate:"eqfield=password"`
}

type DeleteUserRequest struct {
	Username string `json:"username" validate:"required,max=50"`
}

type GetUserByUsernameRequest struct {
	Username string `json:"username" validate:"required,max=50"`
}

type AlamatResponse struct {
	ID        uint   `json:"id"`
	NamaJalan string `json:"nama_jalan"`
	RT        int    `json:"rt"`
	RW        int    `json:"rw"`
	Kelurahan string `json:"kelurahan"`
	Kecamatan string `json:"kecamatan"`
	Kota      string `json:"kota"`
	KodePos   int    `json:"kode_pos"`
}

type UserResponse struct {
	ID                 uint            `json:"id"`
	NamaLengkap        string          `json:"nama_lengkap"`
	Username           string          `json:"username"`
	Email              string          `json:"email"`
	Alamat             *AlamatResponse `json:"alamat"`
	NoHP               string          `json:"no_hp"`
}

type MessageResp struct {
	Message string `json:"message"`
}
