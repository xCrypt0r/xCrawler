import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.net.HttpURLConnection;
import java.net.URL;

public class loadRank { 
	public static int maxPage = 5;
	public static int rankLimit = maxPage * 20;
	public static String[] info = new String[rankLimit];
	
    public static void main(String[] args) {
    	parseRank[] p = new parseRank[maxPage];
    	Thread[] tr = new Thread[maxPage];
    	int threadNumber = maxPage - 1;
    	long start = System.currentTimeMillis();
    	
    	for(int i = 0; i <= threadNumber; i++) {
    		p[i] = new parseRank(i);
    		tr[i] = new Thread(p[i]);
    		tr[i].start();
    	}
    	
    	try {
	    	for(int i = 0; i <= threadNumber; i++) {
	        	tr[i].join();
	    	}
	    	
	        for(int i = 0; i <= rankLimit - 1; i++) {
	        	System.out.println(info[i]);
	        }
	        
	        long end = System.currentTimeMillis();
	        
	        System.out.println("Execution Time : " + (end - start) / 1000.0);
    	} catch(InterruptedException e) {
    		System.out.println(e);
    	}
    }
    
    

	public static class parseRank implements Runnable { 
		int n;
		
	    public parseRank(int n) {
	        this.n = n + 1;
	    }
	
	    public void run() {
	    	try {
	    		URL url = new URL("https://maple.gg/rank/dojang?page=" + n);
	    		HttpURLConnection http = (HttpURLConnection) url.openConnection();
	    		http.setRequestMethod("GET");
	    		
	    		String rank_str = "", server = "", nickname = "", level = "";
	    		int rank = '0';
	    		BufferedReader bfread = new BufferedReader(new InputStreamReader(http.getInputStream()));
	       		String readHTML;
	       		
	    		while ((readHTML = bfread.readLine()) != null) { 
	    			if(readHTML.contains("class=\"text-center align-middle\"")) {
	    				rank_str = readHTML.replaceAll("(^\\s+<th.+\">|</th>)", "");
	    				rank = Integer.parseInt(rank_str);
	    			}
	    			
	    			if(readHTML.contains("<img src=\"https://kr-cdn.maple.gg/images/maplestory/world/ico_world") && readHTML.contains("class=\"text-grape-fruit\"")) {
	    				server = readHTML.replaceAll("(.+alt=\"|\"> <span>.+)", "");
	    				nickname = readHTML.replaceAll("(.+class=\"text-grape-fruit\">|</a>.+)", "");
	    			}
	    			
	    			if(readHTML.contains("span class=\"font-size-14\"")) {
	    				level = readHTML.replaceAll("(.+\">|<.+)", "");
	    			}
	    			
	    			info[rank-1] = rank + "위 [" + server + "] " + nickname + " " + level;
	    		}
	    		
	    		bfread.close();
	    	} catch (IOException e) {
	    		System.out.println(e);
	    	}
	    } 
	}
}
