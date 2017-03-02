#!/bin/bash
# simple test for reproducibility, probably needs major improvements
echo running test-reproducible.sh
set -o errexit
set -o pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && cd .. && pwd )"	# dir!
cd "$DIR" >/dev/null	# work from main gd3 directory
make build
T=`mktemp --tmpdir -d tmp.X'X'X`	# add quotes to avoid matching three X's
cp -a ./gd3 "$T"/gd3.1
make clean
make build
cp -a ./gd3 "$T"/gd3.2

# size comparison test
[ `stat -c '%s' "$T"/gd3.1` -eq `stat -c '%s' "$T"/gd3.2` ] || failures="Size of binary was not reproducible"

# sha1sum test
sha1sum "$T"/gd3.1 > "$T"/gd3.SHA1SUMS.1
sha1sum "$T"/gd3.2 > "$T"/gd3.SHA1SUMS.2
cat "$T"/gd3.SHA1SUMS.1 | sed 's/gd3\.1/gd3\.X/' > "$T"/gd3.SHA1SUMS.1X
cat "$T"/gd3.SHA1SUMS.2 | sed 's/gd3\.2/gd3\.X/' > "$T"/gd3.SHA1SUMS.2X
diff -q "$T"/gd3.SHA1SUMS.1X "$T"/gd3.SHA1SUMS.2X || failures=$( [ -n "${failures}" ] && echo "$failures" ; echo "SHA1SUM of binary was not reproducible" )

# clean up
if [ "$T" != '' ]; then
	rm -rf "$T"
fi
make clean

# display errors
if [[ -n "${failures}" ]]; then
	echo 'FAIL'
	echo 'The following tests failed:'
	echo "${failures}"
	exit 1
fi
echo 'PASS'
