### format-clean-architecture

- Cấu trúc thư mục:
```bash
├── api
│   ├── middleware
│   ├── presenter
│   ├── router
│   │   ├── router.go
│   │   ├── gin.go
├── cmd
│   ├── main.go
├── config
│   ├── config.go
├── controller
│   ├── <entity>.go
├── entity
│   ├── <entity>.go
├── repository
│   ├── mysql
│   ├── postgres
│   │   ├── db.go
│   │   ├── store.go: các hàm thực hiện tương tác qua lại giữa các table
│   │   ├── <entity>.go
├── service
│   ├── interface
│   │   ├── repository
│   │   ├── usecase
│   └── <entity_folder>
│   		└── service.go
```
		
- Các nội dung được yêu cầu hôm trước:
	Yêu cầu 1: api: có thể có tạo nhiều version or dùng các framework khác nhau
	Yêu cầu 2: folder infraction, nếu chỉ chứa repository thì nên tách ra luôn
	Yêu cầu 3: folder usecase: tách thành folder interface(repository + usecase), service
	Yêu cầu 4: main.go: có thể custom nhiều framework
	
	Yêu cầu 2: Ổn, không thấy các vấn đề khi khai báo, sử dụng
	Yêu cầu 3: Em thấy không ổn lắm:
		- Như trong cây thư mục em mô tả ở trên thì ở trong folder interface em không đặt lại 1 folder service nữa
		vì khi import ở file khác thì nó không thể hiện rõ.
		- Còn nếu khai báo là "package interface" thì bị báo lỗi vì tên interface là một khai báo trong go
		- Em có thực hiện triển khai import qua các file khác thì nó không được thống nhất cho lắm giữa các entity
		--- Ý kiến --- Việc chia trong folder service thành các folder entity sẽ ổn hơn về việc khai báo
	Yêu cầu 1, 4: Hiện tại thì em thấy việc sử dụng các framework khác nhau thì em chia trong folder router
	và chỉ cần gọi từ hàm main vài sài
	Việc chia version như anh bảo hôm bữa thì sẽ đổi framwork thì e thấy việc thiết kế router như hiện tại thì vẫn ỏn
	Còn với việc update các luồng xử lý bên ngoài thì có thể tạo các folder version trong folder controller,...
	
THẮC MẮC: việc e sử dụng file store.go trong folder repository để khai báo các hàm tương tác qua lại các table có hợp lý không ạ.

	
	
	
	
	
