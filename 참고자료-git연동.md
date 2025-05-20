# 🌐 Wails 프로젝트 GitHub 연동 가이드

이 문서는 로컬에서 만든 Wails 프로젝트(`tagserver`)를 GitHub와 연동하여 소스코드를 버전 관리하고 공유하는 방법을 정리한 것입니다.

---

## ✅ 1단계: Git 초기화 및 Github 연동

터미널에서 프로젝트 폴더로 이동한 후 Git 저장소를 초기화합니다:

```bash
git init
git add .
git commit -m '최초 커밋: Wails 프로젝트 초기화' 
```

## ✅ 2단계: GitHub에서 새 저장소 만들기
GitHub 접속: https://github.com
오른쪽 상단 ‘+’ 버튼 → New repository
저장소 이름: tagserver 또는 원하는 이름
- README, .gitignore, License 항목 선택하지 않음
[Create repository] 클릭 후 리포지 주소확인

## ✅ 3단계: GitHub 원격 저장소 연결 및 업로드

```bash
git remote add origin https://github.com/AIFarm-team-MISO/tagserver.git
git branch -M main
git push -u origin main
git push

```
## ✅ 4단계: 줄바꿈형식 통일하기 (git에서 자꾸 경고주닊. 필수는 아님 )
.gitattributes 파일 생성후 내용 붙여넣기 
이후 커밋 




