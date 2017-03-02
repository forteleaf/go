Sub 입금확인증()
Dim x As Integer
Dim sort As String
x = ActiveCell.Row

Select Case ActiveSheet.Cells(x, 9).Value
Case "CF"
sort = "■CF □영화 □뮤직비디오 □드라마 □영상물 □정사진 □기타"
Case "영화"
sort = "□CF ■영화 □뮤직비디오 □드라마 □영상물 □정사진 □기타"
Case "드라마"
sort = "□CF □영화 □뮤직비디오 ■드라마 □영상물 □정사진 □기타"
Case "뮤직비디오"
sort = "□CF □영화 ■뮤직비디오 □드라마 □영상물 □정사진 □기타"
Case "정사진"
sort = "□CF □영화 □뮤직비디오 □드라마 □영상물 ■정사진 □기타"
Case "영상물"
sort = "□CF □영화 □뮤직비디오 □드라마 ■영상물 □정사진 □기타"
Case Else
sort = "□CF □영화 □뮤직비디오 □드라마 □영상물 □정사진 ■기타"
End Select

<<<<<<< HEAD
Worksheets("입금확인증").Cells(2, 1) = ActiveSheet.Cells(x, 19).Value '발행번호
Worksheets("입금확인증").Cells(4, 3) = ActiveSheet.Cells(x, 3).Value '제작사
=======
Worksheets("입금확인증").Cells(2, 1) = ActiveSheet.Cells(x, 19).Value '입금확인증 번호
Worksheets("입금확인증").Cells(4.3) = ActiveSheet.Cells(x, 3).Value '제작사
>>>>>>> 765b941b24ff8be327e4ebb68bbb43186d592ebc
Worksheets("입금확인증").Cells(5, 3) = sort '촬영구분
Worksheets("입금확인증").Cells(7, 3) = ActiveSheet.Cells(x, 10).Value '촬영일자
Worksheets("입금확인증").Cells(9, 3) = ActiveSheet.Cells(x, 12).Value '촬영장소
Worksheets("입금확인증").Cells(6, 3) = ActiveSheet.Cells(x, 13).Value '촬영내용
Worksheets("입금확인증").Cells(10, 7) = ActiveSheet.Cells(x, 15).Value '금액(숫자)
Worksheets("입금확인증").Cells(11, 3) = ActiveSheet.Cells(x, 16).Value '입금일자
Worksheets("입금확인증").Cells(12, 3) = ActiveSheet.Cells(x, 20).Value '입금명 현제작사
Worksheets("입금확인증").Select
'Worksheets("입금확인증").PrintOut

End Sub
<<<<<<< HEAD


Sub 기부금영수증()
'일련번호
'기부자성명
'사업자등록증
'주소
'금액
=======
Sub 기부금영수증()
>>>>>>> 765b941b24ff8be327e4ebb68bbb43186d592ebc
x = ActiveCell.Row

Worksheets("기부금영수증").Cells(4, 2) = ActiveSheet.Cells(x, 18).Value '기부금영수증 일련번호
Worksheets("기부금영수증").Cells(8, 4) = ActiveSheet.Cells(x, 3).Value '법인명
<<<<<<< HEAD
Worksheets("기부금영수증").Cells(8, 11) = ActiveSheet.Cells(x, 17).Value '사업자등록번호
Worksheets("기부금영수증").Cells(10, 4) = ActiveSheet.Cells(x, 4).Value '주소
Worksheets("기부금영수증").Cells(27, 6) = ActiveSheet.Cells(x, 16).Value '기부내용 연월일
Worksheets("기부금영수증").Cells(27, 12) = ActiveSheet.Cells(x, 15).Value '금액(숫자)
Worksheets("기부금영수증").Cells(36, 12) = ActiveSheet.Cells(x, 21).Value '대표


=======
Worksheets("기부금영수증").Cells(8, 8) = ActiveSheet.Cells(x, 17).Value '사업자등록번호
Worksheets("기부금영수증").Cells(10, 4) = ActiveSheet.Cells(x, 4).Value '주소
Worksheets("기부금영수증").Cells(27, 12) = ActiveSheet.Cells(x, 15).Value '금액(숫자)
Worksheets("기부금영수증").Cells(36, 12) = ActiveSheet.Cells(x, 21).Value '대표

>>>>>>> 765b941b24ff8be327e4ebb68bbb43186d592ebc
Worksheets("기부금영수증").Select

End Sub
