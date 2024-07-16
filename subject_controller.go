package main

import (
	"fmt"
	"github.com/coredemo/framework"
)

func SubjectAddController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectAddController")
	return nil
}

func SubjectListController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectListController")
	return nil
}

func SubjectDelController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	subjectId, _ := c.ParamInt("id", 0)
	c.SetStatus(200).Json("ok, SubjectGetController:" + fmt.Sprint(subjectId))
	return nil
}

func SubjectNameController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectNameController")
	return nil
}
