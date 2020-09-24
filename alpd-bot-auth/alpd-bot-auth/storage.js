const FAKE_USER_RECORDS = {
    admin: {
	password: 'f15a4214582bbdcfad9bdf01986484b9',
    },
}


exports.findUserByName = (userName) => {
    return FAKE_USER_RECORDS[userName]
}
