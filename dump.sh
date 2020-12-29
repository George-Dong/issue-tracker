#! /bin/sh
cd /tracker
mysqldump -u root -ppassword tracker | sed 's$VALUES ($VALUES\n($g' | sed 's$),($),\n($g' > db.sql