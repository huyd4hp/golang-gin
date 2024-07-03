package schema

type LogOut struct{

}

func (LogOut) TableName() string {
    return "auths"
}