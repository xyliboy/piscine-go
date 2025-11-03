HERO_ID="${HERO_ID:-1}"

curl -s https://platform.zone01.gr/assets/superhero/all.json \
| jq -r --argjson id "$HERO_ID" '.[] | select(.id == $id) | .relatives // empty'
