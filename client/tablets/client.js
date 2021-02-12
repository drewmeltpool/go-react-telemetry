const http = require("../common/http");

const Client = (baseUrl) => (
  (client = http.Client(baseUrl)),
  {
    ListTablets: () => client.get("/tablets"),
    UpdateDevice: (id, Battery, CurrentVideo,DeviceTime) => client.patch(`/tablets`, { id, Battery, CurrentVideo, DeviceTime }),
  }
);

module.exports = { Client };