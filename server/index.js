const grpc = require("@grpc/grpc-js");
const express = require("express");
const protoLoader = require("@grpc/proto-loader");

const app = express();
const PORT = 3000;

const PROTO_PATH = "../siswa/siswa.proto";
const packageDefinition = protoLoader.loadSync(PROTO_PATH);
const siswaProto = grpc.loadPackageDefinition(packageDefinition).siswa;

const siswaData = [
  { id: "1", nama: "John Doe", umur: 25 },
  { id: "2", nama: "Jane Doe", umur: 22 },
];

const server = new grpc.Server();

const serverAddress = "127.0.0.1:50051";

server.addService(siswaProto.SiswaService.service, {
  GetSiswa: (call, callback) => {
    const id = call.request.id;
    const siswa = siswaData.find((s) => s.id === id);
    if (siswa) {
      callback(null, siswa);
    } else {
      callback({
        code: grpc.status.NOT_FOUND,
        details: "Siswa not found",
      });
    }
  },
});

server.bindAsync(
  serverAddress,
  grpc.ServerCredentials.createInsecure(),
  (err, port) => {
    if (err) {
      console.error(`Error binding server: ${err}`);
    } else {
      server.start();
      console.log(`Server is running on ${serverAddress}`);
    }
  }
);

app.listen(PORT, () => {
  console.log(`Express server is running on http://localhost:${PORT}`);
});
