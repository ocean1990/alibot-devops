const PROTO_PATH = __dirname + '/protos/auth.proto'

const grpc = require('grpc')
const protoLoader = require('@grpc/proto-loader')
const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    { keepCase: true, longs: String, enums: String, defaults: true, oneofs: true })
const auth_proto = grpc.loadPackageDefinition(packageDefinition).alpd
const storage = require('./storage')


const port = 9001


function authWithPasswordV1(call, callback) {
    const userInfo = storage.findUserByName(call.request.user)
    const authenticated = userInfo && userInfo.password === call.request.password
    callback(null, {authenticated, msg: ''})
}

function main() {
    const server = new grpc.Server()
    server.addService(auth_proto.AuthService.service, { authWithPasswordV1: authWithPasswordV1 })
    console.log(`start auth server on 0.0.0.0:${port}`)
    server.bind(`0.0.0.0:${port}`, grpc.ServerCredentials.createInsecure())
    server.start()
}

main()
