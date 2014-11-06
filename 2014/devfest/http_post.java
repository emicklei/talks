	String url = "http://www.bol.com/nl/catalog/browse_and_search/plain_list_page.html";
	URL obj = new URL(url);
	HttpsURLConnection con = (HttpsURLConnection) obj.openConnection();
	con.setRequestMethod("POST");
	con.setRequestProperty("Accept-Language", "nl,en;q=0.5");

	String urlParameters = "&Ntt=devfest";
	con.setDoOutput(true);
	DataOutputStream wr = new DataOutputStream(con.getOutputStream());
	wr.writeBytes(urlParameters);
	wr.flush();
	wr.close();

	int responseCode = con.getResponseCode();

	BufferedReader in = new BufferedReader(
	        new InputStreamReader(con.getInputStream()));
	String inputLine;
	StringBuffer response = new StringBuffer();

	while ((inputLine = in.readLine()) != null) {
		response.append(inputLine);
	}
	in.close();
	System.out.println(response.toString());