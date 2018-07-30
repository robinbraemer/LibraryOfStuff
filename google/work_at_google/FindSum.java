/*
 * Copyright (c) RoboFlax. All rights reserved.
 * Use is subject to license terms.
 */
package eu.roboflax.test;

import org.junit.Test;

import java.util.HashSet;
import java.util.Set;

/**
 * Source - https://www.youtube.com/watch?v=XKu_SEDAykw
 * Testing if there is a pair in the list added up equals to the given sum.
 */
public class Algorithm_FindSum {
    
    public static int sum = 8;
    
    public static int[] numbers = {1, 2, 3, 9};
    public static int[] numberz = {1, 2, 4, 6};
    
    @Test
    public void testQuadratic( ) {
        // slow
        System.out.println( "Quadratic" );
        System.out.println( quadratic( numbers ) );
        System.out.println( quadratic( numberz ) );
    }
    
    @Test
    public void testScan( ) {
        // faster, but list must be sorted
        System.out.println( "Scan" );
        System.out.println( scan( numbers ) );
        System.out.println( scan( numberz ) );
    }
    
    @Test
    public void testComplementSearch( ) {
        // fastest and list not must be sorted
        System.out.println( "ComplementSearch" );
        System.out.println( complementSearch( numbers ) );
        System.out.println( complementSearch( numberz ) );
    }
    
    public boolean complementSearch( int[] numbers ) {
        Set<Integer> comp = new HashSet<>(); // complements
        for ( int value : numbers ) {
            if(comp.contains( value ))
                return true;
            comp.add( sum - value );
        }
        return false;
    }
    
    public boolean scan( int[] numbers ) {
        int front = 0;
        int back = numbers.length - 1;
        
        int s;
        while ( front < back ) {
            s = numbers[front] + numbers[back];
            if ( s == sum )
                return true;
            
            if ( s > sum )
                back--;
            else front++;
        }
        return false;
    }
    
    
    public boolean quadratic( int[] numbers ) {
        // iterate all
        for ( int i = 0; i <= numbers.length - 1; i++ ) {
            // test for
            for ( int j = i + 1; j <= numbers.length - 1; j++ ) {
                if ( i == j ) {
                    continue;
                }
                if ( numbers[i] + numbers[j] == sum ) {
                    return true;
                }
            }
        }
        
        return false;
    }
}
