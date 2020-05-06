export class EmailAlreadyExist extends Error {
    constructor(message: string) {
      super(message); 
      this.name = "EmailAlreadyExist";
    }
}