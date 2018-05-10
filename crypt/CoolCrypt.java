
import org.junit.Test;
import sun.misc.BASE64Decoder;
import sun.misc.BASE64Encoder;

import java.io.IOException;
import java.util.Date;
import java.util.Random;

public class TestCrypt {
    
    @Test
    public void test( ) {
        String password = "my1secret2password";
        String encoding = encrypt( password );
        
        System.out.println( "Encrypted string: " + encoding );
        System.out.println( "Decrypted string: " + decrypt( encoding ) );
    }
    
    private static Random rand = new Random( (new Date()).getTime() );
    
    public static String encrypt( String str ) {
        BASE64Encoder encoder = new BASE64Encoder();
        byte[] salt = new byte[8];
        rand.nextBytes( salt );
        return encoder.encode( salt ) + encoder.encode( str.getBytes() );
    }
    
    public static String decrypt( String encStr ) {
        if ( encStr.length() > 12 ) {
            String cipher = encStr.substring( 12 );
            BASE64Decoder decoder = new BASE64Decoder();
            try {
                return new String( decoder.decodeBuffer( cipher ) );
            } catch ( IOException e ) {
                //throw new UnsupportedEncodingException("Tried to decrypt an unsupported encoding");
                return null;
            }
        }
        return null;
    }
}
