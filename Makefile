build:
	go build  ./main.go

run:
	go run ./main.go

serve:
	go run ./main.go serve

install-serve:
	go install cms && cms serve

live-serve:
	HOST="localhost" gin -i --port=8080 --laddr=127.0.0.1 --all run serve

live-test:
	HOST="localhost" gin -i --port=8080 --laddr=127.0.0.1 --all run test

mock:
	mockgen -destination=api/mocks/mock_store.go -package=mocks github.com/ainsleyclark/verbis/api/models AuthRepository,CategoryRepository,FieldsRepository,MediaRepository,OptionsRepository,PostsRepository,RoleRepository,SeoMetaRepository,SessionRepository,SiteRepository,UserRepository

install:
	go install cms