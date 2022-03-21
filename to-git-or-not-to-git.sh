curl -s https://zone01normandie.org/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"Svalbard\"}}){id}}"}' | cut -b 24,25,26
