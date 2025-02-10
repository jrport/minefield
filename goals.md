# GOALS
- [x] Random ip makes request gets user_auth_id
- [ ] faz requisicao no web socket mandando o user_auth_id
    - [ ] caso negativo leva n na cara
    - [ ] eh botado na fila pro matchmaking rolar
- [ ] match anom_id to another anom_id based on avaibality
- each has its own websocket
- [ ] WRAP THE FOLLOWING IN A TIMEOUT CONTEXT
    - on match, group them and give them a canvas with the bombs
    - get their ids (the ones in redis and store them somewhere for the match duration in a persistent manner)
    - signal to the clients we are ready to rumble
    - [ ] get two thumbs up then run
- [ ] client sends a coordinate
- [ ] check what is in their canvas 
- [ ] send response to both

----------------------------------------------------------------

- match anom_id to another anom_id based on a password?

----------------------------------------------------------------
