var sandbox = require('sinon').createSandbox();

import UserRepository from "../../infrastucture/repository/userRepository";
import { CreateUserUseCase } from "./registerUserUseCase";
import  { initializeDBConnection } from "../../infrastucture/database/mongo";

describe("Testing couponService create", async() => {
    const mockDB =  await initializeDBConnection("mock", "mock");
    const mockRepo = new UserRepository(mockDB, "mock_collection");
    const createUserUseCase = new CreateUserUseCase(mockRepo);
    beforeEach(function () {
        // stub out the `hello` method
        sandbox.stub(initializeDBConnection);
        sandbox.stub(UserRepository);
        mo
    });
    it("It should call CouponRepo create once and return coupon", async () => {
        const userData = {
            name: "shahin",
            age: 19,
            gender: "Male",
        }
        mockCreateUser.mockReturnValueOnce(userData);

        const res = await createUserUseCase.createUser(userData);

        expect(mockCreateUser).toHaveBeenCalledTimes(1);
        expect(res).toHaveProperty("adult");
    });
});