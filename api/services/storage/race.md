WARNING: DATA RACE
Write at 0x00c000410858 by goroutine 48:
github.com/verbiscms/verbis/api/services/storage.(*StorageTestSuite).BeforeTest()
/Users/ainsley/Desktop/Reddico/verbis/verbis/api/services/storage/storage_test.go:45 +0x66
github.com/stretchr/testify/suite.Run.func1()
/Users/ainsley/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:151 +0x8a5
testing.tRunner()
/usr/local/go/src/testing/testing.go:1194 +0x202

Previous write at 0x00c000410858 by goroutine 47:
bytes.(*Buffer).grow()
/usr/local/go/src/bytes/buffer.go:144 +0x27a
bytes.(*Buffer).Write()
/usr/local/go/src/bytes/buffer.go:172 +0x184
github.com/sirupsen/logrus.(*Entry).write()
/Users/ainsley/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/entry.go:286 +0x2de
github.com/sirupsen/logrus.(*Entry).log()
/Users/ainsley/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/entry.go:251 +0x384
github.com/sirupsen/logrus.(*Entry).Log()
/Users/ainsley/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/entry.go:293 +0xc4
github.com/sirupsen/logrus.(*Entry).Error()
/Users/ainsley/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/entry.go:322 +0x36b
github.com/verbiscms/verbis/api/services/storage.(*MigrationInfo).fail()
/Users/ainsley/Desktop/Reddico/verbis/verbis/api/services/storage/migrate.go:80 +0x2e8
github.com/verbiscms/verbis/api/services/storage.(*Storage).migrateBackground()
/Users/ainsley/Desktop/Reddico/verbis/verbis/api/services/storage/migrate.go:246 +0x9c7

Goroutine 48 (running) created at:
testing.(*T).Run()
/usr/local/go/src/testing/testing.go:1239 +0x5d7
github.com/stretchr/testify/suite.runTests()
/Users/ainsley/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:203 +0xf7
github.com/stretchr/testify/suite.Run()
/Users/ainsley/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:176 +0x944
github.com/verbiscms/verbis/api/services/storage.TestStorage()
/Users/ainsley/Desktop/Reddico/verbis/verbis/api/services/storage/storage_test.go:39 +0x5e
testing.tRunner()
/usr/local/go/src/testing/testing.go:1194 +0x202

Goroutine 47 (finished) created at:
github.com/verbiscms/verbis/api/services/storage.(*Storage).processMigration()
/Users/ainsley/Desktop/Reddico/verbis/verbis/api/services/storage/migrate.go:218 +0x55d
==================
==================
WARNING: DATA RACE
Write at 0x00c000410870 by goroutine 48:
github.com/verbiscms/verbis/api/services/storage.(*StorageTestSuite).BeforeTest()
/Users/ainsley/Desktop/Reddico/verbis/verbis/api/services/storage/storage_test.go:45 +0x66
github.com/stretchr/testify/suite.Run.func1()
/Users/ainsley/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:151 +0x8a5
testing.tRunner()
/usr/local/go/src/testing/testing.go:1194 +0x202

Previous write at 0x00c000410870 by goroutine 45:
bytes.(*Buffer).grow()
/usr/local/go/src/bytes/buffer.go:147 +0x2bb
bytes.(*Buffer).Write()
/usr/local/go/src/bytes/buffer.go:172 +0x184
github.com/sirupsen/logrus.(*Entry).write()
/Users/ainsley/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/entry.go:286 +0x2de
github.com/sirupsen/logrus.(*Entry).log()
/Users/ainsley/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/entry.go:251 +0x384
github.com/sirupsen/logrus.(*Entry).Log()
/Users/ainsley/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/entry.go:293 +0xc4
github.com/sirupsen/logrus.(*Logger).Log()
/Users/ainsley/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/logger.go:198 +0xa4
github.com/sirupsen/logrus.(*Logger).Info()
/Users/ainsley/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/logger.go:220 +0x692
github.com/verbiscms/verbis/api/logger.Info()
/Users/ainsley/Desktop/Reddico/verbis/verbis/api/logger/logger.go:78 +0x64c
github.com/verbiscms/verbis/api/services/storage.(*Storage).processMigration()
/Users/ainsley/Desktop/Reddico/verbis/verbis/api/services/storage/migrate.go:223 +0x588

Goroutine 48 (running) created at:
testing.(*T).Run()
/usr/local/go/src/testing/testing.go:1239 +0x5d7
github.com/stretchr/testify/suite.runTests()
/Users/ainsley/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:203 +0xf7
github.com/stretchr/testify/suite.Run()
/Users/ainsley/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:176 +0x944
github.com/verbiscms/verbis/api/services/storage.TestStorage()
/Users/ainsley/Desktop/Reddico/verbis/verbis/api/services/storage/storage_test.go:39 +0x5e
testing.tRunner()
/usr/local/go/src/testing/testing.go:1194 +0x202

Goroutine 45 (finished) created at:
github.com/verbiscms/verbis/api/services/storage.(*Storage).Migrate()
/Users/ainsley/Desktop/Reddico/verbis/verbis/api/services/storage/migrate.go:163 +0xd67
github.com/verbiscms/verbis/api/services/storage.(*StorageTestSuite).TestStorage_Migrate.func7()
/Users/ainsley/Desktop/Reddico/verbis/verbis/api/services/storage/migrate_test.go:136 +0x2a6
github.com/stretchr/testify/suite.(*Suite).Run.func1()
/Users/ainsley/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:77 +0x172
testing.tRunner()
/usr/local/go/src/testing/testing.go:1194 +0x202
==================
==================
WARNING: DATA RACE
Write at 0x00c000410878 by goroutine 48:
github.com/verbiscms/verbis/api/services/storage.(*StorageTestSuite).BeforeTest()
/Users/ainsley/Desktop/Reddico/verbis/verbis/api/services/storage/storage_test.go:45 +0x66
github.com/stretchr/testify/suite.Run.func1()
/Users/ainsley/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:151 +0x8a5
testing.tRunner()
/usr/local/go/src/testing/testing.go:1194 +0x202

Previous write at 0x00c000410878 by goroutine 45:
bytes.(*Buffer).Write()
/usr/local/go/src/bytes/buffer.go:169 +0x44
github.com/sirupsen/logrus.(*Entry).write()
/Users/ainsley/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/entry.go:286 +0x2de
github.com/sirupsen/logrus.(*Entry).log()
/Users/ainsley/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/entry.go:251 +0x384
github.com/sirupsen/logrus.(*Entry).Log()
/Users/ainsley/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/entry.go:293 +0xc4
github.com/sirupsen/logrus.(*Logger).Log()
/Users/ainsley/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/logger.go:198 +0xa4
github.com/sirupsen/logrus.(*Logger).Info()
/Users/ainsley/go/pkg/mod/github.com/sirupsen/logrus@v1.8.1/logger.go:220 +0x79e
github.com/verbiscms/verbis/api/logger.Info()
/Users/ainsley/Desktop/Reddico/verbis/verbis/api/logger/logger.go:78 +0x758
github.com/verbiscms/verbis/api/services/storage.(*Storage).processMigration()
/Users/ainsley/Desktop/Reddico/verbis/verbis/api/services/storage/migrate.go:224 +0x693

Goroutine 48 (running) created at:
testing.(*T).Run()
/usr/local/go/src/testing/testing.go:1239 +0x5d7
github.com/stretchr/testify/suite.runTests()
/Users/ainsley/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:203 +0xf7
github.com/stretchr/testify/suite.Run()
/Users/ainsley/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:176 +0x944
github.com/verbiscms/verbis/api/services/storage.TestStorage()
/Users/ainsley/Desktop/Reddico/verbis/verbis/api/services/storage/storage_test.go:39 +0x5e
testing.tRunner()
/usr/local/go/src/testing/testing.go:1194 +0x202

Goroutine 45 (finished) created at:
github.com/verbiscms/verbis/api/services/storage.(*Storage).Migrate()
/Users/ainsley/Desktop/Reddico/verbis/verbis/api/services/storage/migrate.go:163 +0xd67
github.com/verbiscms/verbis/api/services/storage.(*StorageTestSuite).TestStorage_Migrate.func7()
/Users/ainsley/Desktop/Reddico/verbis/verbis/api/services/storage/migrate_test.go:136 +0x2a6
github.com/stretchr/testify/suite.(*Suite).Run.func1()
/Users/ainsley/go/pkg/mod/github.com/stretchr/testify@v1.7.0/suite/suite.go:77 +0x172
testing.tRunner()
/usr/local/go/src/testing/testing.go:1194 +0x202
==================
--- FAIL: TestStorage (0.52s)
--- FAIL: TestStorage/TestStorage_MigrateBackground (0.01s)
testing.go:1093: race detected during execution of test
testing.go:1093: race detected during execution of test
FAIL
exit status 1
FAIL    github.com/verbiscms/verbis/api/services/storage        0.761s
ainsley@Ainsleys-iMac-Pro storage % go test --race
