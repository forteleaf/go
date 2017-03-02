package com.xmlrpc.tistory;
import java.net.URL;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import org.apache.xmlrpc.client.XmlRpcClient;
import org.apache.xmlrpc.client.XmlRpcClientConfigImpl;


public class XmlRpcTistoryBlog {

    static final String API_URL = "http://regexr.tistory.com/api";


    static final String API_ID = "1111";
    static final String USER_ID = "user@gmail.com";
    static final String API_PASSWORD = "2222";

    public static void main(String[] args) {
        // TODO Auto-generated method stub

        try {

            XmlRpcClientConfigImpl config = new XmlRpcClientConfigImpl();
            config.setServerURL(new URL(API_URL));

            Map<String, Object> contents = new HashMap<String, Object>();

            String categories[] = new String[] {"UNIX"};
            contents.put("categories", categories); // 카테고리 배열
            contents.put("title", "여기에 제목을 쓰고"); // 제목
            contents.put("description", "여기에 내용을 쓰세요."); // 내용
            contents.put("mt_keywords", "태그1, 태그2, 태그3"); // 태크 콤마로 구분한다.

            List<Object> params = new ArrayList<Object>();

            // 여기에는 API정보에 설정된 ID가 들어간다
            params.add(API_ID);

            // 사용자 계정 즉 이메일이 들어간다.
            params.add(USER_ID);

            // API 암호 또는 유저암호 둘다 가능하다.
            params.add(API_PASSWORD);

            // 블로그 컨텐츠
            params.add(contents);

            // 공개여부 true이면 공개, false면 비공개
            params.add(new Boolean(true));

            XmlRpcClient client = new XmlRpcClient();
            client.setConfig(config);

            String rsString = (String) client.execute("metaWeblog.newPost", params);

            System.out.println(rsString);

        }catch(Exception e) {
            e.printStackTrace();
        }

    }
}
