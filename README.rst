Welcome to go-semver's documentation!
=====================================


go-semver is a golang package provides a simple support for `Semantic Versioning <http://semver.org/>`_.


.. image:: https://secure.travis-ci.org/Lispython/go-semver.png
	   :target: https://secure.travis-ci.org/Lispython/go-semver


INSTALLATION
------------

To use ``go-semverzz execute next command::

  go install github.com/Lispython/go-semver


It's install ``semver`` package to ``GOROOT`` path.



USAGE
-----

To use library in your packages import it first.

.. sourcecode:: go

   package your_package_name

   import (
      "github.com/Lispython/go-semver"
      "fmt"
    )

   var (
        VERSION *semver.Version = &Version{1, 10, 8}
    )

    func main(){
       fmt.Println(VERSION.ToString())
	}




CONTRIBUTE
----------

We need you help.

#. Check for open issues or open a fresh issue to start a discussion around a feature idea or a bug.
   There is a Contributor Friendly tag for issues that should be ideal for people who are not very familiar with the codebase yet.
#. Fork `the repository`_ on Github to start making your changes to the **develop** branch (or branch off of it).
#. Write a test which shows that the bug was fixed or that the feature works as expected.
#. Send a pull request and bug the maintainer until it gets merged and published.

.. _`the repository`: https://github.com/Lispython/go-semver/
