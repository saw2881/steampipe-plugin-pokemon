# Table: pokemon_type

Types are properties for Pokémon and their moves. Each type has three properties: which types of Pokémon it is super effective against, which types of Pokémon it is not very effective against, and which types of Pokémon it is completely ineffective against.

## Examples

### Basic info

```sql
select   
  name, 
  id 
from  
  pokemon_type
```

### Get all the pokemons whose names starts with 'fly'

```sql
select   
  name, 
    id 
from  
  pokemon_type 
where 
  name like  'fly%'
```

### Get the pokemon with an id of 17

```sql
select   
  name, 
  id 
from  
  pokemon_type 
where 
id = 17
```
