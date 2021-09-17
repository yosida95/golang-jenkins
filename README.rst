golang-jenkins
==============

-----
About
-----
This is a API client of Jenkins API written in Go.

-----
Usage
-----

``go get -u github.com/garudaen/golang-jenkins``

``import "github.com/garudaen/golang-jenkins"``

Configure authentication and create an instance of the client:

.. code-block:: go

   auth := &gojenkins.Auth{
      Username: "[jenkins user name]",
      ApiToken: "[jenkins API token]",
   }
   jenkins := gojenkins.NewJenkins(auth, "[jenkins instance base url]")

Make calls against the desired resources:

.. code-block:: go

   job, err := jenkins.GetJob("[job name]")

-------
License
-------
golang-jenkins is licensed under the MIT LICENSE.
See `./LICENSE <./LICENSE>`_.
