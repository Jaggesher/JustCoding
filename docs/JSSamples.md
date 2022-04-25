# Some JS Sample

- Inheritance With Function
``` js
function Identity(name, age){
  this.name = name
  this.age = age
}
Identity.prototype.revealIdentity = function(){
  console.log(`Name: ${this.name}, Age: ${this.age}`)
}

function Elf(name, age, weapon){
  this.weapon = weapon
  Identity.call(this, name, age)
}

Elf.prototype = Object.create(Identity.prototype);
Elf.prototype.constructor = Elf;
Elf.prototype.attack = function() {
  console.log(`${this.name}, attack with ${this.weapon}`)
};

function Human(name, age, weapon, lover){
  this.lover = lover
  Elf.call(this, name, age, weapon)
}

Human.prototype = Object.create(Elf.prototype);
Human.prototype.constructor = Human;
Human.prototype.loves = function(){
  console.log(`${this.name} Loves ${this.lover}`)
}

const peter = new Identity('Peter', 12);
peter.revealIdentity();

const robin = new Elf('robin', 32, 'stone');
robin.attack();

const jack = new Human('Jack', 42, 'Ak47', 'Rose');
jack.revealIdentity();
jack.attack();
jack.loves();
```