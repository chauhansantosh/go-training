module example.com/mainpkg

go 1.18

replace example.com/firstassessment => ../firstassessment

replace example.com/secondassessment => ../secondassessment

replace example.com/thirdassessment => ../thirdassessment

require (
	example.com/firstassessment v0.0.0-00010101000000-000000000000
	example.com/fourthassessment v0.0.0-00010101000000-000000000000
	example.com/secondassessment v0.0.0-00010101000000-000000000000
	example.com/thirdassessment v0.0.0-00010101000000-000000000000
)

require example.com/fifthassessment v0.0.0-00010101000000-000000000000

replace example.com/fourthassessment => ../fourthassessment

replace example.com/fifthassessment => ../fifthassessment
