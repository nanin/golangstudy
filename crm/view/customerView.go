package view

import (
	"fmt"
	"test/crm/model"
	"test/crm/service"
)

type customerView struct {
	service *service.CustomerService
}

func CrmRun() {

	customerView := customerView{
		service: service.CreateCustomerService(),
	}
	customerView.InitMenu()
}

func (view *customerView) InitMenu() {

	isBreak := false
	for {

		fmt.Println("======欢迎使用CRM客户关系系统======")
		fmt.Println("          1、新增客户")
		fmt.Println("          2、修改客户")
		fmt.Println("          3、删除客户")
		fmt.Println("          4、查询列表")
		fmt.Println("          5、退   出")
		fmt.Println("请选择菜单：")
		var menuIndex int
		fmt.Scanln(&menuIndex)
		switch menuIndex {
		case 1:
			view.addCustomer()
		case 2:
			view.updateCustomer()
		case 3:
			view.deleteCustomer()
		case 4:
			view.list()
		case 5:
			isBreak = true
		}
		if isBreak {
			return
		}
	}

}

func (view *customerView) updateCustomer() {
	var code int
	fmt.Print("请输入编号:")
	fmt.Scanln(&code)
	index, info := view.service.GetIndex(code)
	fmt.Println("编号\t姓名\t电话\t地址")
	fmt.Println(info.DetilStringInfo())

	fmt.Println("===========开始修改客户=============")
	var tel int
	fmt.Print("请输入电话")
	fmt.Scanln(&tel)
	var address string
	fmt.Print("请输入地址")
	fmt.Scanln(&address)
	info.Update(tel, address)
	view.service.UpdateName(index, info)
	fmt.Println("===========修改客户成功=============")
}

func (view *customerView) deleteCustomer() {
	var code int
	fmt.Print("请输入编号:")
	fmt.Scanln(&code)

	fmt.Println("===========开始删除客户=============")
	ind, _ := view.service.GetIndex(code)
	if ind != -1 {
		view.service.Delete(ind)
	}
	fmt.Println("===========删除客户成功=============")
}

func (view *customerView) addCustomer() {
	fmt.Println("===========开始新增客户=============")
	var name string
	fmt.Print("请输入姓名")
	fmt.Scanln(&name)
	var tel int
	fmt.Print("请输入电话")
	fmt.Scanln(&tel)
	var address string
	fmt.Print("请输入地址")
	fmt.Scanln(&address)
	cus := model.CreateCustomerWithoutId(name, tel, address)
	view.service.Add(*cus)

	fmt.Println("===========新增客户成功=============")
}

func (view *customerView) list() {
	custoerms := view.service.List()
	fmt.Println("===========开始查询列表=============")
	if len(custoerms) <= 0 {
		fmt.Println("暂时没有客户")
	} else {
		fmt.Println("编号\t姓名\t电话\t地址")
		for i := 0; i < len(custoerms); i++ {
			fmt.Println(custoerms[i].DetilStringInfo())
		}
	}

	// fmt.Println(view.service.List())
	fmt.Println("===========列表查询成功=============")
}
