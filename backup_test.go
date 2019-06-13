package sessiongo

import (
	"log"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_backup(t *testing.T) {
	s := NewSessionsStorage()
	session := NewSession("test_id")
	session.SetData("test")
	s.Add(session)
	Convey("test init", t, func() {
		session, ok := s.Get("test_id")
		So(ok, ShouldBeTrue)

		data := session.GetData()
		So(data.(string), ShouldEqual, "test")

		Convey("test backup", func() {
			err := s.BackUp()
			So(err, ShouldBeNil)

			file, err := os.Open(backup)
			So(err, ShouldBeNil)
			defer file.Close()
		})
		Convey("test read backup", func() {
			session := NewSession("test_readbackup")
			session.SetData("backup")
			s.Add(session)

			session, ok := s.Get("test_readbackup")
			So(ok, ShouldBeTrue)

			data := session.GetData()
			log.Println(data.(string))
			So(data.(string), ShouldEqual, "backup")

			err := s.ReadBackup()
			So(err, ShouldBeNil)

			_, ok = s.Get("test_backup")
			So(ok, ShouldBeFalse)

			session, ok = s.Get("test_id")
			So(ok, ShouldBeTrue)

			data = session.GetData()
			So(data.(string), ShouldEqual, "test")

		})
	})

}
