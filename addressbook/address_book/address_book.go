package address_book

/*
*  结构体包含 联系人姓名
*			联系人电话
* 			联系人邮箱
*			最后拨打电话时间
*/
type contacts struct {
	name string
	phone string
	email string
	lastTimeContacts string
}

type AddressBook struct {
	contactsArray []contacts
	size int
}

// 最大联系人存贮人数
const CAPACITY int = 1000

// 通讯录初始化

func (addressBook *AddressBook) AddressBookInit() {
	addressBook.contactsArray = make([]contacts, CAPACITY)
	addressBook.size = 0
}

// 添加联系人
//	多参数类型一致 只用在最后写
//
func (ct *contacts) addContact(name, phone, email string) {

}

func (ct *contacts) removeContact() {

}

func (ct *contacts) updateContact() {

}

func (ct *contacts) searchContact(name string) {

}