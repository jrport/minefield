# GOALS
- [x] Random ip makes request gets anom_id
- [x] Bota na fila
- [x] match anom_id to another anom_id based on avaibality
- create one websocket for each and sync the state of a message to both

----------------------------------------------------------------

- match anom_id to another anom_id based on a password?


----------------------------------------------------------------

client: chama no websocket mandando o id no head auth
server: busca pelo id na queue
server: marca esse cara e bota numa outra fila marcando q ta ready
server: manda um wait pelo ws
server: quando tiver dois >= 2 users ready
server: associa ambos a um mesmo buffer
