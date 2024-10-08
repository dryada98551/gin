package libmodel

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

// Account 對應DB中的 account
type Account struct {
	ID         uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID     uuid.UUID `gorm:"column:user_id;type:uuid;unique;not null" json:"user_id" validate:"required"`
	UserNo     string    `gorm:"column:user_no;type:VARCHAR(9);unique;not null" json:"user_no" validate:"required"`
	WebOwnerID uuid.UUID `gorm:"column:webowner_id;type:uuid;not null" json:"webowner_id" validate:"required"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;not null;default:(CURRENT_TIMESTAMP AT TIME ZONE 'UTC')"" json:"created_at"`
}

func ValidateAccount(table *Account) error {
	validate := validator.New()
	return validate.Struct(table)
}

// UserInfo 對應DB中的 user_info
// Status   | 凍結 = 1 | 刪除 = 99
type UserInfo struct {
	ID               uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID           uuid.UUID `gorm:"column:user_id;type:uuid;unique;not null" json:"user_id" validate:"required"`
	FirstName        string    `gorm:"column:first_name;type:varchar(30)" json:"first_name"`
	LastName         string    `gorm:"column:last_name;type:varchar(30)" json:"last_name"`
	NickName         string    `gorm:"column:nick_name;type:varchar(100)" json:"nick_name"`
	Gender           string    `gorm:"column:gender;type:varchar(100)" json:"gender"`
	Birthday         string    `gorm:"column:birthday;type:varchar(30)" json:"birthday"`
	Nation           string    `gorm:"column:nation;type:varchar(100)" json:"nation"`
	BornPlace        string    `gorm:"column:born_place;type:varchar(100)" json:"born_place"`
	CurrentAddress   string    `gorm:"column:current_address;type:varchar(100)" json:"current_address"`
	PermanentAddress string    `gorm:"column:permanent_address;type:varchar(100)" json:"permanent_address"`
	Email            string    `gorm:"column:email;type:varchar(100)" json:"email"`
	JobType          string    `gorm:"column:job_type;type:varchar(30)" json:"job_type"`
	IncomeOrigin     string    `gorm:"column:income_origin;type:varchar(100)" json:"income_origin"`
	Phone            string    `gorm:"column:phone;type:varchar(30)" json:"phone"`
	Status           int16     `gorm:"column:status;type:int2" json:"status"`
	CreatedAt        time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func ValidateUserInfo(table *UserInfo) error {
	validate := validator.New()
	return validate.Struct(table)
}

// UserLoginStatus 對應DB中的 user_login_status
type UserLoginStatus struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID `gorm:"column:user_id;type:uuid;unique;not null" json:"user_id" validate:"required"`
	JwtToken  string    `gorm:"column:jwt_token;type:VARCHAR(1024);not null" json:"jwt_token" validate:"required"`
	IsLogin   bool      `gorm:"column:is_login;type:boolean;not null;default:false" json:"is_login" validate:"required"`
	LoginAt   time.Time `gorm:"column:login_at;type:timestamp;not null;" json:"login_at" validate:"required"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;default:(CURRENT_TIMESTAMP AT TIME ZONE 'UTC')"" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;not null;default:(CURRENT_TIMESTAMP AT TIME ZONE 'UTC')"" json:"updated_at"`
}

func ValidateUserLoginStatus(table *UserLoginStatus) error {
	validate := validator.New()
	return validate.Struct(table)
}

// UserSocial 對應DB中的 user_social
type UserSocial struct {
	ID         uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID     uuid.UUID `gorm:"column:user_id;type:uuid;unique;not null" json:"user_id" validate:"required"`
	UserToken  uuid.UUID `gorm:"column:user_token;type:uuid;unique;not null" json:"user_token" validate:"required"`
	SocialType string    `gorm:"column:social_type;type:VARCHAR(50);not null" json:"social_type" validate:"required"`
	SocialID   string    `gorm:"column:social_id;type:varchar(50);" json:"social_id"`
	FirstName  string    `gorm:"column:first_name;type:varchar(30);" json:"first_name"`
	LastName   string    `gorm:"column:last_name;type:varchar(30);" json:"last_name"`
	Email      string    `gorm:"column:email;type:varchar(100);" json:"email"`
	CreatedAt  time.Time `gorm:"column:created_at;type:timestamp;not null;default:(CURRENT_TIMESTAMP AT TIME ZONE 'UTC')"" json:"created_at"`
}

func ValidateUserSocial(table *UserSocial) error {
	validate := validator.New()
	return validate.Struct(table)
}

// UserPhone 對應DB中的 user_phone
type UserPhone struct {
	ID              uuid.UUID `gorm:"column:id;type:uuid;primary_key;" json:"id"`
	UserID          uuid.UUID `gorm:"column:user_id;type:uuid;unique;not null" json:"user_id" validate:"required"`
	UserPhoneNumber string    `gorm:"column:user_phone_number;type:VARCHAR(30);unique;not null" json:"user_phone_number" validate:"required"`
	CreatedAt       time.Time `gorm:"column:created_at;type:timestamp;not null;default:(CURRENT_TIMESTAMP AT TIME ZONE 'UTC')"" json:"created_at"`
}

func ValidateUserPhone(table *UserPhone) error {
	validate := validator.New()
	return validate.Struct(table)
}

// UserLoginLog 對應DB中的 user_login_log
type UserLoginLog struct {
	ID           uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID       uuid.UUID `gorm:"column:user_id;type:uuid;not null" json:"user_id" validate:"required"`
	LoginType    int16     `gorm:"column:login_type;type:int2;not null" json:"login_type" validate:"min=0"`
	IP           string    `gorm:"column:ip;type:varchar(30)" json:"ip"`
	Country      string    `gorm:"column:country;type:varchar(50)" json:"country"`
	Browser      string    `gorm:"column:browser;type:varchar(50)" json:"browser"`
	DeviceType   string    `gorm:"column:device_type;type:varchar(50)" json:"device_type"`
	DeviceVendor string    `gorm:"column:device_vendor;type:varchar(50)" json:"device_vendor"`
	OS           string    `gorm:"column:os;type:varchar(50)" json:"os"`
	Location     string    `gorm:"column:location;type:varchar(50)" json:"location"`
	UserAgent    string    `gorm:"column:user_agent;type:varchar(255)" json:"user_agent"`
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func ValidateUserLoginLog(table *UserLoginLog) error {
	validate := validator.New()
	return validate.Struct(table)
}

// WebownerSetting 對應DB中的 webowner_setting
type WebownerSetting struct {
	ID             uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	WebownerID     uuid.UUID `gorm:"column:webowner_id;type:uuid;not null" json:"webowner_id" validate:"required"`
	Type           int16     `gorm:"column:type;type:int2;not null" json:"type" validate:"min=0"`
	TypeName       string    `gorm:"column:type_name;type:varchar(30);not null" json:"type_name" validate:"required"`
	MiddleItem     int16     `gorm:"column:middle_item;type:int2;not null" json:"middle_item" validate:"min=0"`
	MiddleItemName string    `gorm:"column:middle_item_name;type:varchar(30);not null" json:"middle_item_name" validate:"required"`
	Item           int16     `gorm:"column:item;type:int2;not null" json:"item" validate:"min=0"`
	ItemName       string    `gorm:"column:item_name;type:varchar(30);not null" json:"item_name" validate:"required"`
	Val            string    `gorm:"column:val;type:varchar(30);not null" json:"val" validate:"required"`
	Mark           string    `gorm:"column:mark;type:varchar(100)" json:"mark"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func ValidateWebownerSetting(table *WebownerSetting) error {
	validate := validator.New()
	return validate.Struct(table)
}

// UserKYC 對應DB中的 user_kyc
// Type   | 會員資料 = 1 | 手機認證 = 2 | 身分驗證 = 3 | 生物辨識 = 4
// Status | 待審核 = 1 | 待補件 = 2 | 已通過 = 3
type UserKYC struct {
	ID         uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID     uuid.UUID `gorm:"column:user_id;type:uuid;not null" json:"user_id" validate:"required"`
	UserKYCID  uuid.UUID `gorm:"column:user_kyc_id;type:uuid;unique;not null" json:"user_kyc_id" validate:"required"`
	Type       int16     `gorm:"column:type;type:int2;not null" json:"type" validate:"min=0"`
	Status     int16     `gorm:"column:status;type:int2;not null" json:"status" validate:"min=0"`
	CompleteAt time.Time `gorm:"column:complete_at;not null" json:"complete_at" validate:"required"`
	CreatedAt  time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func ValidateUserKYC(table *UserKYC) error {
	validate := validator.New()
	return validate.Struct(table)
}

// UserKYCReason 對應DB中的 user_kyc_reason
type UserKYCReason struct {
	ID          uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserKYCID   uuid.UUID `gorm:"column:user_kyc_id;type:uuid;unique;not null" json:"user_kyc_id" validate:"required"`
	KYCOptionID uuid.UUID `gorm:"column:kyc_option_id;type:uuid;not null" json:"kyc_option_id" validate:"required"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func ValidateUserKYCReason(table *UserKYCReason) error {
	validate := validator.New()
	return validate.Struct(table)
}

// UserKYCReasonOption 對應DB中的 user_kyc_reason_option
type UserKYCReasonOption struct {
	ID          uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	KYCOptionID uuid.UUID `gorm:"column:kyc_option_id;type:uuid;unique;not null" json:"kyc_option_id" validate:"required"`
	Type        int16     `gorm:"column:type;type:int2;not null" json:"type" validate:"min=0"`
	ItemName    string    `gorm:"column:item_name;type:varchar(100);not null" json:"item_name" validate:"required"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func ValidateUserKYCReasonOption(table *UserKYCReasonOption) error {
	validate := validator.New()
	return validate.Struct(table)
}

// UserBalance 對應DB中的 user_balance
type UserBalance struct {
	ID             uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID         uuid.UUID `gorm:"column:user_id;type:uuid;unique;not null" json:"user_id" validate:"required"`
	Balance        float64   `gorm:"column:balance;type:double precision;not null" json:"balance" validate:"min=0"`
	DisableBalance float64   `gorm:"column:disable_balance;type:double precision;not null" json:"disable_balance" validate:"min=0"`
	TotalPayment   float64   `gorm:"column:total_payment;type:double precision" json:"total_payment"`
	Level          string    `gorm:"column:level;type:varchar(30);not null" json:"level" validate:"required"`
	NextLevel      string    `gorm:"column:next_level;type:varchar(30)" json:"next_level"`
	BetAmount      float64   `gorm:"column:bet_amount;type:double precision;not null" json:"bet_amount" validate:"min=0"`
	NeedAmount     float64   `gorm:"column:need_amount;type:double precision" json:"need_amount"`
	Ranking        int       `gorm:"column:ranking;type:int4" json:"ranking"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func ValidateUserBalance(table *UserBalance) error {
	validate := validator.New()
	return validate.Struct(table)
}

// UserFund 對應DB中的 user_fund
type UserFund struct {
	ID          uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	OrderID     uuid.UUID `gorm:"column:order_id;type:uuid;unique;not null" json:"order_id" validate:"required"`
	UserID      uuid.UUID `gorm:"column:user_id;type:uuid;not null" json:"user_id" validate:"required"`
	Platform    string    `gorm:"column:platform;type:varchar(50);not null" json:"platform" validate:"required"`
	AccountType int16     `gorm:"column:account_type;type:int2;not null" json:"account_type" validate:"min=0"`
	FundType    int16     `gorm:"column:fund_type;type:int2;not null" json:"fund_type" validate:"min=0"`
	PayType     int16     `gorm:"column:pay_type;type:int2;not null" json:"pay_type" validate:"min=0"`
	PayMoney    float64   `gorm:"column:pay_money;type:double precision;not null" json:"pay_money" validate:"min=0"`
	BeforeMoney float64   `gorm:"column:before_money;type:double precision;not null" json:"before_money" validate:"min=0"`
	AfterMoney  float64   `gorm:"column:after_money;type:double precision;not null" json:"after_money" validate:"min=0"`
	Remark      string    `gorm:"column:remark;type:varchar(255)" json:"remark"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func ValidateUserFund(table *UserFund) error {
	validate := validator.New()
	return validate.Struct(table)
}

// UserWithdraw 對應DB中的 user_withdraw
type UserWithdraw struct {
	ID             uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	OrderID        uuid.UUID `gorm:"column:order_id;type:uuid;unique;not null" json:"order_id" validate:"required"`
	UserID         uuid.UUID `gorm:"column:user_id;type:uuid;not null" json:"user_id" validate:"required"`
	Platform       string    `gorm:"column:platform;type:varchar(50);not null" json:"platform" validate:"required"`
	BankCode       string    `gorm:"column:bank_code;type:varchar(30);not null" json:"bank_code" validate:"required"`
	BankCardNumber string    `gorm:"column:bank_card_number;type:varchar(255);not null" json:"bank_card_number" validate:"required"`
	WithdrawType   int16     `gorm:"column:withdraw_type;type:int2;not null" json:"withdraw_type" validate:"min=0"`
	WithdrawStatus int16     `gorm:"column:withdraw_status;type:int2;not null" json:"withdraw_status" validate:"min=0"`
	RequestMoney   float64   `gorm:"column:request_money;type:double precision;not null" json:"request_money" validate:"min=0"`
	SendMoney      float64   `gorm:"column:send_money;type:double precision;not null" json:"send_money" validate:"min=0"`
	Remark         string    `gorm:"column:remark;type:varchar(255)" json:"remark"`
	ResponseMoney  float64   `gorm:"column:response_money;type:double precision" json:"response_money"`
	ResponseMsg    string    `gorm:"column:response_msg;type:varchar(255)" json:"response_msg"`
	Fee            float64   `gorm:"column:fee;type:double precision;not null" json:"fee" validate:"min=0"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func ValidateUserWithdraw(table *UserWithdraw) error {
	validate := validator.New()
	return validate.Struct(table)
}

// UserCharge 對應DB中的 user_charge
type UserCharge struct {
	ID            uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	OrderID       uuid.UUID `gorm:"column:order_id;type:uuid;unique;not null" json:"order_id" validate:"required"`
	UserID        uuid.UUID `gorm:"column:user_id;type:uuid;not null" json:"user_id" validate:"required"`
	Platform      string    `gorm:"column:platform;type:varchar(50);not null" json:"platform" validate:"required"`
	ChannelCode   string    `gorm:"column:channel_code;type:varchar(50);not null" json:"channel_code" validate:"required"`
	ChargeStatus  int16     `gorm:"column:charge_status;type:int2;not null" json:"charge_status" validate:"min=0"`
	RequestMoney  float64   `gorm:"column:request_money;type:double precision;not null" json:"request_money" validate:"min=0"`
	SendMoney     float64   `gorm:"column:send_money;type:double precision;not null" json:"send_money" validate:"min=0"`
	Remark        string    `gorm:"column:remark;type:varchar(255)" json:"remark"`
	ResponseMoney float64   `gorm:"column:response_money;type:double precision" json:"response_money"`
	ResponseMsg   string    `gorm:"column:response_msg;type:varchar(255)" json:"response_msg"`
	Fee           float64   `gorm:"column:fee;type:double precision;not null" json:"fee" validate:"min=0"`
	CreatedAt     time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func ValidateUserCharge(table *UserCharge) error {
	validate := validator.New()
	return validate.Struct(table)
}

// ChargeChannel 對應DB中的 charge_channel
type ChargeChannel struct {
	ID          uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	ChannelCode string    `gorm:"column:channel_code;type:varchar(50);unique;not null" json:"channel_code" validate:"required"`
	Platform    string    `gorm:"column:platform;type:varchar(50)" json:"platform"`
	ChannelName string    `gorm:"column:channel_name;type:varchar(100)" json:"channel_name"`
	IsEnable    bool      `gorm:"column:is_enable;type:boolean;not null" json:"is_enable" validate:"required"`
	PaymentType int16     `gorm:"column:payment_type;type:int2;not null" json:"payment_type" validate:"min=0"`
	MinMoney    int64     `gorm:"column:min_money;type:int8" json:"min_money"`
	MaxMoney    int64     `gorm:"column:max_money;type:int8" json:"max_money"`
	Fee         float64   `gorm:"column:fee;type:double precision" json:"fee"`
	ReturnURL   string    `gorm:"column:return_url;type:varchar(255)" json:"return_url"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func ValidateChargeChannel(table *ChargeChannel) error {
	validate := validator.New()
	return validate.Struct(table)
}

// GameStatus 對應DB中的 game_status
type GameStatus struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID `gorm:"column:user_id;type:uuid;unique;not null" json:"user_id" validate:"required"`
	Status    int16     `gorm:"column:status;type:int2;not null" json:"status" validate:"min=0"`
	Factory   string    `gorm:"column:factory;type:varchar(50)" json:"factory"`
	GameID    string    `gorm:"column:game_id;type:varchar(30)" json:"game_id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func ValidateGameStatus(table *GameStatus) error {
	validate := validator.New()
	return validate.Struct(table)
}

// GameAccount 對應DB中的 game_account
type GameAccount struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID    uuid.UUID `gorm:"column:user_id;type:uuid;unique;not null" json:"user_id" validate:"required"`
	AccountID string    `gorm:"column:account_id;type:varchar(100);unique;not null" json:"account_id" validate:"required"`
	Factory   string    `gorm:"column:factory;type:varchar(50)" json:"factory"`
	GameID    string    `gorm:"column:game_id;type:varchar(30)" json:"game_id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func ValidateGameAccount(table *GameAccount) error {
	validate := validator.New()
	return validate.Struct(table)
}

// GameTransferOrder 對應DB中的 game_transfer_order
type GameTransferOrder struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	OrderID   uuid.UUID `gorm:"column:order_id;type:uuid;unique;not null" json:"order_id" validate:"required"`
	UserID    uuid.UUID `gorm:"column:user_id;type:uuid;not null" json:"user_id" validate:"required"`
	AccountID string    `gorm:"column:account_id;type:varchar(100);not null" json:"account_id" validate:"required"`
	Factory   string    `gorm:"column:factory;type:varchar(50)" json:"factory"`
	GameID    string    `gorm:"column:game_id;type:varchar(30)" json:"game_id"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func ValidateGameTransferOrder(table *GameTransferOrder) error {
	validate := validator.New()
	return validate.Struct(table)
}

// SettingTask 對應DB中的 setting_task
type SettingTask struct {
	ID         uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	WebOwnerID uuid.UUID `gorm:"column:webowner_id;type:uuid;not null" json:"webowner_id" validate:"required"`
	TaskID     uuid.UUID `gorm:"column:task_id;type:uuid;not null" json:"task_id" validate:"required"`
	TaskName   string    `gorm:"column:task_name;type:varchar(30);not null" json:"task_name" validate:"required"`
	TaskStatus int16     `gorm:"column:task_status;type:int2;not null" json:"task_status" validate:"min=0"`
	CreatedAt  time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func ValidateSettingTask(table *SettingTask) error {
	validate := validator.New()
	return validate.Struct(table)
}

// TaskStatus 對應DB中的 task_status
type TaskStatus struct {
	ID         uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID     uuid.UUID `gorm:"column:user_id;type:uuid;not null" json:"user_id" validate:"required"`
	WebOwnerID uuid.UUID `gorm:"column:webowner_id;type:uuid;not null" json:"webowner_id" validate:"required"`
	TaskID     uuid.UUID `gorm:"column:task_id;type:uuid;not null" json:"task_id" validate:"required"`
	TaskName   string    `gorm:"column:task_name;type:varchar(30);not null" json:"task_name" validate:"required"`
	TaskStatus int16     `gorm:"column:task_status;type:int2;not null" json:"task_status" validate:"min=0"`
	CreatedAt  time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func ValidateTaskStatus(table *TaskStatus) error {
	validate := validator.New()
	return validate.Struct(table)
}

// TaskList 對應DB中的 task_list
type TaskList struct {
	ID        uuid.UUID `gorm:"column:id;type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	TaskID    uuid.UUID `gorm:"column:task_id;type:uuid;unique;not null" json:"task_id" validate:"required"`
	TaskName  string    `gorm:"column:task_name;type:varchar(30);not null" json:"task_name" validate:"required"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func ValidateTaskList(table *TaskList) error {
	validate := validator.New()
	return validate.Struct(table)
}
