import { ListUserUseCase } from "./listUserUseCase"
import { Gender, User } from "../../entity/user";

// TODO:-> mocks should be separate folder
let mockUserRepository = {
  createUser: jest.fn(),
  getUserByEmail: jest.fn(),
  listUser: jest.fn(),
};


describe("Testing ListUserUseCase", () => {
    it("it should return a user", async() => {
        mockUserRepository.createUser.mockClear();
        mockUserRepository.getUserByEmail.mockClear();
        mockUserRepository.listUser.mockClear();

        const user = User.FromUser({
          ID: "sdfsdf",
          age: 2,
          email: "sdfsdf",
          gender: Gender.Male,
          name: "sdfsd",
          password: "234343",
          phone: "2343434"
        });
        mockUserRepository.listUser.mockResolvedValueOnce([user]);
        
        const listUserUseCase = new ListUserUseCase(mockUserRepository);
  
        const users = await listUserUseCase.execute(0, 20);
        
        expect(users.length).toBe(1);
        expect(mockUserRepository.listUser).toHaveBeenCalled();       
    });
});