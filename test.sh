SUCCESS=0
TOTAL=100

echo "Début du test de disponibilité..."
for i in $(seq 1 $TOTAL); do
  HTTP_CODE=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8081/healthz --max-time 1)
  echo "Requête $i : $HTTP_CODE"
  if [ "$HTTP_CODE" -eq 200 ]; then
    SUCCESS=$((SUCCESS + 1))
  fi
done

PERCENTAGE=$((SUCCESS * 100 / TOTAL))
echo "Résultat : $PERCENTAGE% des requêtes ont réussi !"