Pidofmain=$(pidof main)
Pidofswagger=$(pidof swagger)

if [ $Pidofmain ]; then
  kill $Pidofmain
  echo "$Pidofmain killed"
else
  echo "main never start"
fi

if [ $Pidofswagger ]; then
  kill $Pidofswagger
  echo "$Pidofswagger killed"
else
  echo "swagger never start"
fi