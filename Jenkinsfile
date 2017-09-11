pipeline {
  agent any
  stages {
    stage('') {
      steps {
        parallel(
          "1": {
            echo 'testing new pipelines'
            
          },
          "2": {
            echo 'something in parallel?'
            
          }
        )
      }
    }
    stage('sh') {
      steps {
        sh 'sh "ls -la"'
      }
    }
  }
}