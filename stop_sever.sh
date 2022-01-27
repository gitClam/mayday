Pidofmain=$(pidof main)


if [ $Pidofmain ]; then
  kill $Pidofmain
  echo "$Pidofmain killed"
else
  echo "main never start"
fi
