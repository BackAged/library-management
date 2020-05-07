import { RegisterUserUseCase } from "./registerUserUseCase"
import { mocked } from 'ts-jest/utils'
import { Gender, User } from "../../entity/user";
import { EmailAlreadyExist } from "./errors";



let mockUserRepository = {
  createUser: jest.fn(),
  getUserByEmail: jest.fn(),
};
let mockJsonWebTokenManager = {
  generateJSONWebTOken: jest.fn(),
}

describe("Testing RegisterUserUseCase", () => {
    it("it should register a user", async() => {
      const user = User.FromUser({
        ID: "sdfsdf",
        age: 2,
        email: "sdfsdf",
        gender: Gender.Male,
        name: "sdfsd",
        password: "234343",
        phone: "2343434"
      })
      mockUserRepository.createUser.mockResolvedValueOnce(user);
      
      const registerUseCase = new RegisterUserUseCase(mockUserRepository, mockJsonWebTokenManager);

      await registerUseCase.execute({
        age: 2,
        email: "sdfsdf",
        gender: Gender.Male,
        name: "sdfsd",
        password: "234343",
        phone: "2343434"
      });

      expect(mockUserRepository.createUser).toHaveBeenCalled();
      expect(mockJsonWebTokenManager.generateJSONWebTOken).toHaveBeenCalled();
    });

    it("it should throw when a user email already exist", async() => {
      const user = User.FromUser({
        ID: "sdfsdf",
        age: 2,
        email: "sdfsdf",
        gender: Gender.Male,
        name: "sdfsd",
        password: "234343",
        phone: "2343434"
      })
      mockUserRepository.getUserByEmail.mockResolvedValueOnce(user);
      
      const registerUseCase = new RegisterUserUseCase(mockUserRepository, mockJsonWebTokenManager);

      try {
        await registerUseCase.execute({
          age: 2,
          email: "sdfsdf",
          gender: Gender.Male,
          name: "sdfsd",
          password: "234343",
          phone: "2343434"
        });
      } catch(e) {
        expect(e).toBeInstanceOf(EmailAlreadyExist)
      }
      

      expect(mockUserRepository.createUser).toHaveBeenCalled();
      expect(mockJsonWebTokenManager.generateJSONWebTOken).toHaveBeenCalled();
    });
});