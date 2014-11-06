	InputStream fis = null;
	try {
	    fis = new FileInputStream("2014/devfest/quotes.txt");
	} catch (FileNotFoundException e) {
	    e.printStackTrace();
	}
	InputStreamReader ir = new InputStreamReader(fis, Charset.forName("UTF-8"))
	BufferedReader br = new BufferedReader(ir);
	String line;
	try {
	    while ((line = br.readLine()) != null) {
	        System.out.println(line);
	    }
	    br.close();
	} catch (IOException e) {
	    e.printStackTrace();
	}