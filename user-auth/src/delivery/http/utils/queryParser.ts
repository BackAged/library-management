export const skipLimitParser = (query: any): { skip: number; limit: number } => {
    const skip = query.skip ? Number(query.skip) : 0;
    const limit = query.limit ? Number(query.limit) > 25 ? 25 : Number(query.limit) : 25;
    return { skip, limit };
};
