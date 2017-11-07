conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
if err != nil {
	log.Fatal("Dial failed:", err)
}
client := NewClockServiceClient(conn)
resp, err := client.GetTime(context.Background(), new(GetTimeRequest))

fmt.Println(resp.FormattedTime)
