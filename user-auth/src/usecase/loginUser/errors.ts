
export class UserNotFound extends Error {
    constructor(message: string) {
      super(message); 
      this.name = "UserNotFound";
    }
}


export class PasswardMisMatch extends Error {
    constructor(message: string) {
      super(message); 
      this.name = "PasswardMisMatch";
    }
}