/*
 * Copyright 2016-2017 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package example;

import java.util.logging.Logger;

/**
 * @author Dave Syer
 *
 */
public class Handler {

    private final static Logger logger = Logger.getLogger(Handler.class.getName());

    public static native void start();

    public static native void writeResponse(String input);

    public static native String readRequest();

    public static void main(String[] args) {
        logger.info("Starting " + System.getProperty("java.library.path"));
        System.loadLibrary("Hello");
        logger.info("Loaded Hello Golang lib, invoking Go for listening on Lambda Handler");
        start();
        while (true) {
            String request = readRequest();
            String response = request.toUpperCase();
            writeResponse(response);
        }
    }

}
