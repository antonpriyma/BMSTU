
import java.io.File;
import java.io.*;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;




public class Sync {
    private static ArrayList<String> tree =  new ArrayList<>();
    private static ArrayList<Path> D = new ArrayList<>();
    private static ArrayList<Path> S = new ArrayList<>();
    private static ArrayList<Path> deleteList = new ArrayList<>();
    private static ArrayList<Path> copyList = new ArrayList<>();
    private static Path path1;

    private static boolean  sameContent(Path file1, Path file2) throws IOException {
        final long size = Files.size(file1);
        if (size != Files.size(file2))
            return false;

            return Arrays.equals(Files.readAllBytes(file1), Files.readAllBytes(file2));
    }

    private static boolean contain(Path path, ArrayList<Path> f){
        boolean flag=false;
        for (int j=0;j<f.size();j++){
            path1 = f.get(j);
            if (path1.equals(path)){
                flag=true;
                break;
            }
        }
        return flag;
    }

    private static void readPath(String s, ArrayList<Path> f){
        File f1  = new File(s);
        String[] f2 = f1.list();
        for (int i=0;i<f2.length;i++){
            String dir = f2[i];
            File help = new File(s+"/"+dir);
            if (help.isDirectory()){
                readPath(help.toString(),f);
            }
            else {
                String k = help.toString();
                k=k.replaceAll(k.substring(0,1),"");
                f.add(help.toPath().subpath(1,help.toPath().getNameCount()));
            }
        }
    }


    public static void main(String[] argc) throws IOException {
        String test = argc[0];
        String semple = argc[1];
        readPath(test,S);
        readPath(semple,D);

        for(int i=0;i<D.size();i++){
            boolean flag=false;
            Path path = D.get(i);
            for (int j=0;j<S.size();j++){
                path1 = S.get(j);
                if (path1.equals(path)){
                    flag=true;
                    break;
                }
            }
            if (!flag){
                deleteList.add(D.get(i));
            }
        }

        for (int i=0;i<S.size();i++){
            boolean flag=false;
            File buffer = new File(semple+"/"+ S.get(i).toString());
            if (contain(S.get(i),D)){
                File help = new File(test+"/"+path1.toString());
                if (buffer.lastModified()==help.lastModified()) {
                    continue;
                }
                else if (!sameContent(help.toPath(),buffer.toPath())){
                    deleteList.add(S.get(i));
                    copyList.add(S.get(i));
                }
            }
            else{
                copyList.add(S.get(i));
            }
        }
        Collections.sort(deleteList);
        Collections.sort(copyList);
        if (deleteList.isEmpty() && copyList.isEmpty()){
            System.out.println("IDENTICAL");
            System.exit(0);
        }
        for (int i=0;i<deleteList.size();i++){
            System.out.println("DELETE "+deleteList.get(i));
        }

        for (int i=0;i<copyList.size();i++){
            System.out.println("COPY "+copyList.get(i));
        }


    }

}

