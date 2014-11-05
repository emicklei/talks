    import java.util.concurrent.ExecutorService;
	import java.util.concurrent.Executors;
	import java.util.concurrent.TimeUnit;
		
	// ...	
		
		ExecutorService exec = Executors.newCachedThreadPool();
        exec.submit(new Runnable() {
            
            @Override
            public void run() {
                System.out.println("some time later");               
            }
        });
		
        try {
            exec.awaitTermination(10, TimeUnit.SECONDS);
        } catch (InterruptedException e) {            
            e.printStackTrace();
        }
		
	// ...			